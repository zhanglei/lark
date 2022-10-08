package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lark/domain/po"
	"lark/domain/repo"
	"lark/examples/config"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmongo"
	"lark/pkg/common/xredis"
	"lark/pkg/entity"
)

func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	var (
		i              int64
		message        *po.Message
		messageHotRepo repo.MessageHotRepository
		messages       []*po.Message
		err            error
	)
	for i = 0; i < 10; i++ {
		message = &po.Message{
			SrvMsgId:       i,
			CliMsgId:       1,
			RootId:         0,
			ParentId:       0,
			UpperMessageId: 0,
			SenderId:       1,
			ReceiverId:     1,
			SenderPlatform: 1,
			ChatId:         1,
			ChatType:       1,
			SeqId:          i,
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
			if err.(mongo.WriteException).WriteErrors[0].Code == 11000 {
				err = nil
			}
		}
	}

	w := entity.NewMongoWhere()
	w.SetLimit(int64(10))
	w.SetSort("seq_id", true)
	w.SetFilter("chat_type", 1)
	w.SetFilter("chat_id", 1)
	w.SetFilter("seq_id", bson.M{"$gt": 5})
	if messages, err = messageHotRepo.Messages(w); err != nil {
		xlog.Warn(err.Error())
	}
	if len(messages) > 0 {

	}

	u := entity.NewMongoUpdate()
	u.SetFilter("chat_id", 1)
	u.SetFilter("srv_msg_id", 1)
	u.Set("root_id", 10)
	if err = messageHotRepo.Update(u); err != nil {
		xlog.Warn(err.Error())
	}
}
