package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type ChatInviteRepository interface {
	NewChatInvite(req *po.ChatInvite) (err error)
	TxNewChatInviteList(tx *gorm.DB, list []*po.ChatInvite) (err error)
	TxUpdateChatInvite(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	TxChatInvite(tx *gorm.DB, w *entity.MysqlWhere) (invite *po.ChatInvite, err error)
	ChatInvite(w *entity.MysqlWhere) (invite *po.ChatInvite, err error)
	ChatInviteList(w *entity.MysqlWhere) (list []*po.ChatInvite, err error)
}

type chatInviteRepository struct {
}

func NewChatInviteRepository() ChatInviteRepository {
	return &chatInviteRepository{}
}

func (r *chatInviteRepository) NewChatInvite(req *po.ChatInvite) (err error) {
	db := xmysql.GetDB()
	err = db.Create(req).Error
	return
}

func (r *chatInviteRepository) TxNewChatInviteList(tx *gorm.DB, list []*po.ChatInvite) (err error) {
	err = tx.Create(list).Error
	return
}

func (r *chatInviteRepository) TxUpdateChatInvite(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.ChatInvite{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatInviteRepository) TxChatInvite(tx *gorm.DB, w *entity.MysqlWhere) (invite *po.ChatInvite, err error) {
	invite = new(po.ChatInvite)
	err = tx.Where(w.Query, w.Args...).Find(invite).Error
	return
}

func (r *chatInviteRepository) ChatInvite(w *entity.MysqlWhere) (invite *po.ChatInvite, err error) {
	invite = new(po.ChatInvite)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(invite).Error
	return
}

func (r *chatInviteRepository) ChatInviteList(w *entity.MysqlWhere) (list []*po.ChatInvite, err error) {
	list = make([]*po.ChatInvite, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Find(&list).Error
	return
}
