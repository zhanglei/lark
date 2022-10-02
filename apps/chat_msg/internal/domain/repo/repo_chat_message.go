package repo

import (
	"lark/pkg/entity"
)

type ChatMessageRepository interface {
	HistoryMessages(w *entity.MysqlWhere) (list []*entity.Message, err error)
	HotMessages(w *entity.MongoWhere) (list []*entity.Message, err error)
}

type chatMessageRepository struct {
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}
