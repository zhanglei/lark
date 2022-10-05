package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type ChatRepository interface {
	TxCreate(tx *gorm.DB, chat *po.Chat) (err error)
	Chat(w *entity.MysqlWhere) (chat *po.Chat, err error)
}

type chatRepository struct {
}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

func (r *chatRepository) TxCreate(tx *gorm.DB, chat *po.Chat) (err error) {
	err = tx.Create(chat).Error
	return
}

func (r *chatRepository) Chat(w *entity.MysqlWhere) (chat *po.Chat, err error) {
	chat = new(po.Chat)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(chat).Error
	return
}
