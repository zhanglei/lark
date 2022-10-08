package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lark/domain/po"
	"lark/domain/repo"
	"lark/examples/config"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmongo"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	var (
		message        *po.Message
		messageHotRepo repo.MessageHotRepository
		err            error
	)
	message = &po.Message{
		SrvMsgId:       1,
		CliMsgId:       1,
		RootId:         0,
		ParentId:       0,
		UpperMessageId: 0,
		SenderId:       1,
		ReceiverId:     1,
		SenderPlatform: 1,
		ChatId:         1,
		ChatType:       1,
		SeqId:          1,
		MsgFrom:        1,
		MsgType:        1,
		Body:           "Test Message",
		Status:         1,
		SentTs:         1,
		SrvTs:          1,
		UpdatedTs:      1,
		DeletedTs:      0,
	}
	// 消息入库
	messageHotRepo = repo.NewMessageHotRepository()
	if err = messageHotRepo.Create(message); err != nil {
		xlog.Warn(err.Error())
		if err.(*mongo.WriteError).Code == 11000 {
			err = nil
		}
		return
	}
}
