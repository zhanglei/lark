package service

import (
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/proto"
	"lark/domain/po"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/utils"
)

func (s *messageHotService) MessageHandler(msg []byte, msgKey string) (err error) {
	var (
		req     = new(pb_mq.InboxMessage)
		message = new(po.Message)
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, ERROR_MSG_HOT_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	// 消息入库
	copier.Copy(message, req.Msg)
	message.Body = utils.MsgBodyToStr(req.Msg.MsgType, req.Msg.Body)
	message.UpdatedTs = utils.NowMilli()
	if err = s.messageHotRepo.Create(message); err != nil {
		xlog.Warn(err.Error())
		if err.(*mongo.WriteError).Code == constant.ERROR_CODE_MONGOL_DUPLICATE_ENTRY {
			err = nil
		}
		return
	}
	return
}
