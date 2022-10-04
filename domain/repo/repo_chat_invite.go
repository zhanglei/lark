package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type ChatInviteRepository interface {
	RequestCreate(req *po.ChatRequest) (err error)
	TxRequestUpdate(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	TxRequest(tx *gorm.DB, w *entity.MysqlWhere) (request *po.ChatRequest, err error)
	RequestList(w *entity.MysqlWhere) (list []po.ChatRequest, err error)
	TxChatUsersCreate(tx *gorm.DB, users []*po.ChatUser) (err error)
}

type chatInviteRepository struct {
}

func NewChatInviteRepository() ChatInviteRepository {
	return &chatInviteRepository{}
}

func (r *chatInviteRepository) RequestCreate(req *po.ChatRequest) (err error) {
	db := xmysql.GetDB()
	err = db.Create(req).Error
	return
}

func (r *chatInviteRepository) TxRequestUpdate(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.ChatRequest{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatInviteRepository) TxRequest(tx *gorm.DB, w *entity.MysqlWhere) (request *po.ChatRequest, err error) {
	err = tx.Where(w.Query, w.Args...).Find(request).Error
	return
}

func (r *chatInviteRepository) RequestList(w *entity.MysqlWhere) (list []po.ChatRequest, err error) {
	list = make([]po.ChatRequest, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Find(&list).Error
	return
}

func (r *chatInviteRepository) TxChatUsersCreate(tx *gorm.DB, users []*po.ChatUser) (err error) {
	err = tx.Create(users).Error
	return
}
