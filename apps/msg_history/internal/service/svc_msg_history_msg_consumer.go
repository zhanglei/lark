package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *messageHistoryService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_NEW:
		err = s.SaveMessage(msg)
		return
	case constant.CONST_MSG_KEY_RECALL:
		err = s.MessageRecall(msg)
		return
	default:
		return
	}
}

func (s *messageHistoryService) SaveMessage(msg []byte) (err error) {
	var (
		req     = new(pb_mq.InboxMessage)
		message = new(po.Message)
		jsonStr string
		key     string
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	// 消息入库
	copier.Copy(message, req.Msg)
	message.Body = utils.MsgBodyToStr(req.Msg.MsgType, req.Msg.Body)
	if err = s.chatMessageRepo.Create(message); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_INSERT_MESSAGE_FAILED, ERROR_MSG_HISTORY_INSERT_MESSAGE_FAILED, err.Error())
		if err.(*mysql.MySQLError).Number == constant.ERROR_CODE_MYSQL_DUPLICATE_ENTRY {
			err = nil
		}
		return
	}
	// 消息缓存
	jsonStr, _ = utils.Marshal(message)
	key = constant.RK_MSG_CACHE + utils.Int64ToStr(message.ChatId) + ":" + utils.Int64ToStr(message.SeqId)
	if err = xredis.Set(key, jsonStr, constant.CONST_DURATION_MSG_CACHE_SECOND); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_CHCHE_MESSAGE_FAILED, ERROR_MSG_HISTORY_CHCHE_MESSAGE_FAILED, err.Error())
		// 避免再次消费
		err = nil
		return
	}
	return
}

func (s *messageHistoryService) MessageRecall(msg []byte) (err error) {
	var (
		req = new(pb_msg.MessageOperationReq)
		u   = entity.NewMysqlUpdate()
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HISTORY_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	u.SetFilter("srv_msg_id=?", req.SrvMsgId)
	u.SetFilter("sender_id=?", req.SenderId)
	u.Set("status", pb_enum.MSG_OPERATION_RECALL)
	err = s.chatMessageRepo.UpdateMessage(u)
	if err != nil {
		xlog.Warn(ERROR_CODE_MSG_HISTORY_UPDATE_VALUE_FAILED, ERROR_MSG_HISTORY_UPDATE_VALUE_FAILED, err.Error())
	}
	// TODO:缓存更新

	return
}
