package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
)

type MessageHistoryRepository interface {
	Create(message *po.Message) (err error)
}

type messageHistoryRepository struct {
}

func NewMessageHistoryRepository() MessageHistoryRepository {
	return &messageHistoryRepository{}
}

func (r *messageHistoryRepository) Create(message *po.Message) (err error) {
	db := xmysql.GetDB()
	return db.Create(message).Error
}
