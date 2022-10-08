package repo

import (
	"gorm.io/gorm"
	"lark/domain/do"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberRepository interface {
	TxCreate(tx *gorm.DB, chatMember *po.ChatMember) (err error)
	ChatMemberUidList(w *entity.MysqlWhere) (list []int64, err error)
	ChatMemberList(w *entity.MysqlWhere) (list []*do.ChatMemberInfo, err error)
	ChatMemberSetting(w *entity.MysqlWhere) (member *po.ChatMember, err error)
	PushMemberList(w *entity.MysqlWhere) (list []*pb_chat_member.PushMember, err error)
	PushMember(w *entity.MysqlWhere) (conf *pb_chat_member.PushMember, err error)
	ChatMember(w *entity.MysqlWhere) (member *po.ChatMember, err error)
	UpdateChatMember(u *entity.MysqlUpdate) (err error)
	TxUpdateChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	ChatMemberBasicInfoList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberBasicInfo, err error)
}

type chatMemberRepository struct {
}

func NewChatMemberRepository() ChatMemberRepository {
	return &chatMemberRepository{}
}

func (r *chatMemberRepository) TxCreate(tx *gorm.DB, chatMember *po.ChatMember) (err error) {
	err = tx.Create(chatMember).Error
	return
}

func (r *chatMemberRepository) ChatMemberUidList(w *entity.MysqlWhere) (list []int64, err error) {
	list = make([]int64, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(w.Query, w.Args...).Limit(w.Limit).Pluck("uid", &list).Error
	return
}

func (r *chatMemberRepository) ChatMemberList(w *entity.MysqlWhere) (list []*do.ChatMemberInfo, err error) {
	list = make([]*do.ChatMemberInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).
		Select("chat_id,uid,mute,platform,server_id").
		Where(w.Query, w.Args...).
		Limit(w.Limit).Find(&list).Error
	return
}

func (r *chatMemberRepository) ChatMemberSetting(w *entity.MysqlWhere) (member *po.ChatMember, err error) {
	member = new(po.ChatMember)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(w.Query, w.Args...).Find(&member).Error
	return
}

func (r *chatMemberRepository) PushMemberList(w *entity.MysqlWhere) (list []*pb_chat_member.PushMember, err error) {
	list = make([]*pb_chat_member.PushMember, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("uid,mute,platform,server_id").Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *chatMemberRepository) PushMember(w *entity.MysqlWhere) (conf *pb_chat_member.PushMember, err error) {
	conf = new(pb_chat_member.PushMember)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("chat_id,uid,mute,platform,server_id").Where(w.Query, w.Args...).Find(&conf).Error
	return
}

func (r *chatMemberRepository) ChatMember(w *entity.MysqlWhere) (member *po.ChatMember, err error) {
	member = new(po.ChatMember)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(w.Query, w.Args...).Find(&member).Error
	return
}

func (r *chatMemberRepository) UpdateChatMember(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatMemberRepository) TxUpdateChatMember(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatMemberRepository) ChatMemberBasicInfoList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberBasicInfo, err error) {
	list = make([]*pb_chat_member.ChatMemberBasicInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(po.ChatMember{}).Select("uid,display_name,avatar_key").Where(w.Query, w.Args...).
		Limit(w.Limit).Find(&list).Error
	return
}
