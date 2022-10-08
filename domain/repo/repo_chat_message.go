package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type ChatMessageRepository interface {
	Create(message *po.Message) (err error)
	HistoryMessages(w *entity.MysqlWhere) (list []*po.Message, err error)
}

type chatMessageRepository struct {
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}

func (r *chatMessageRepository) Create(message *po.Message) (err error) {
	db := xmysql.GetDB()
	return db.Create(message).Error
}

func (r *chatMessageRepository) HistoryMessages(w *entity.MysqlWhere) (list []*po.Message, err error) {
	list = make([]*po.Message, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Order(w.Sort).Find(&list).Error
	return
}
