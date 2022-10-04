package repo

import (
	"lark/domain/po"
	"lark/pkg/entity"
)

type ChatMessageRepository interface {
	HistoryMessages(w *entity.MysqlWhere) (list []*po.Message, err error)
	HotMessages(w *entity.MongoWhere) (list []*po.Message, err error)
}

type chatMessageRepository struct {
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}
