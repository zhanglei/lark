package repo

import (
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type MessageHistoryRepository interface {
	Create(message *entity.Message) (err error)
}

type messageHistoryRepository struct {
}

func NewMessageHistoryRepository() MessageHistoryRepository {
	return &messageHistoryRepository{}
}

func (r *messageHistoryRepository) Create(message *entity.Message) (err error) {
	db := xmysql.GetDB()
	return db.Create(message).Error
}
