package service

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/utils"
)

func (s *messageHistoryService) MessageHandler(msg []byte, msgKey string) (err error) {
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
