package repos

import (
	"lark/domain/pos"
	"lark/pkg/common/xmysql"
)

type MessageHistoryRepository interface {
	Create(message *pos.Message) (err error)
}

type messageHistoryRepository struct {
}

func NewMessageHistoryRepository() MessageHistoryRepository {
	return &messageHistoryRepository{}
}

func (r *messageHistoryRepository) Create(message *pos.Message) (err error) {
	db := xmysql.GetDB()
	return db.Create(message).Error
}
