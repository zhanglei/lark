package repos

import (
	"lark/domain/pos"
	"lark/pkg/entity"
)

type ChatMessageRepository interface {
	HistoryMessages(w *entity.MysqlWhere) (list []*pos.Message, err error)
	HotMessages(w *entity.MongoWhere) (list []*pos.Message, err error)
}

type chatMessageRepository struct {
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}
