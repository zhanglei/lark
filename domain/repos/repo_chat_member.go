package repos

import (
	"lark/domain/pos"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberRepository interface {
	ChatMemberUidList(w *entity.MysqlWhere) (list []int64, err error)
	ChatMemberList(w *entity.MysqlWhere) (list []*pos.ChatMember, err error)
	ChatMemberSetting(w *entity.MysqlWhere) (member *pos.ChatMember, err error)
	ChatMemberPushConfigList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberPushConfig, err error)
	ChatMemberPushConfig(w *entity.MysqlWhere) (conf *pb_chat_member.ChatMemberPushConfig, err error)
	ChatMemberInfo(w *entity.MysqlWhere) (member *pos.ChatMember, err error)
	UpdateChatMember(u *entity.MysqlUpdate) (err error)
	ChatMemberBasicInfoList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberBasicInfo, err error)
}

type chatMemberRepository struct {
}

func NewChatMemberRepository() ChatMemberRepository {
	return &chatMemberRepository{}
}

func (r *chatMemberRepository) ChatMemberUidList(w *entity.MysqlWhere) (list []int64, err error) {
	list = make([]int64, 0)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Where(w.Query, w.Args...).Limit(w.Limit).Pluck("uid", &list).Error
	return
}

func (r *chatMemberRepository) ChatMemberList(w *entity.MysqlWhere) (list []*pos.ChatMember, err error) {
	list = make([]*pos.ChatMember, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).
		Limit(w.Limit).Find(&list).Error
	return
}

func (r *chatMemberRepository) ChatMemberSetting(w *entity.MysqlWhere) (member *pos.ChatMember, err error) {
	member = new(pos.ChatMember)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Where(w.Query, w.Args...).Find(&member).Error
	return
}

func (r *chatMemberRepository) ChatMemberPushConfigList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberPushConfig, err error) {
	list = make([]*pb_chat_member.ChatMemberPushConfig, 0)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Select("chat_id,uid,mute,platform,server_id").Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *chatMemberRepository) ChatMemberPushConfig(w *entity.MysqlWhere) (conf *pb_chat_member.ChatMemberPushConfig, err error) {
	conf = new(pb_chat_member.ChatMemberPushConfig)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Select("chat_id,uid,mute,platform,server_id").Where(w.Query, w.Args...).Find(&conf).Error
	return
}

func (r *chatMemberRepository) ChatMemberInfo(w *entity.MysqlWhere) (member *pos.ChatMember, err error) {
	member = new(pos.ChatMember)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Where(w.Query, w.Args...).Find(&member).Error
	return
}

func (r *chatMemberRepository) UpdateChatMember(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

func (r *chatMemberRepository) ChatMemberBasicInfoList(w *entity.MysqlWhere) (list []*pb_chat_member.ChatMemberBasicInfo, err error) {
	list = make([]*pb_chat_member.ChatMemberBasicInfo, 0)
	db := xmysql.GetDB()
	err = db.Model(pos.ChatMember{}).Select("uid,display_name,avatar_key").Where(w.Query, w.Args...).
		Limit(w.Limit).Find(&list).Error
	return
}
