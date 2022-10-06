package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type AvatarRepository interface {
	Avatar(w *entity.MysqlWhere) (avatar *po.Avatar, err error)
	AvatarList(w *entity.MysqlWhere) (avatars []*po.Avatar, err error)
	TxSaveAvatar(tx *gorm.DB, avatar *po.Avatar) (err error)
}

type avatarRepository struct {
}

func NewAvatarRepository() AvatarRepository {
	return &avatarRepository{}
}

func (r *avatarRepository) Avatar(w *entity.MysqlWhere) (avatar *po.Avatar, err error) {
	avatar = new(po.Avatar)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(avatar).Error
	return
}

func (r *avatarRepository) AvatarList(w *entity.MysqlWhere) (avatars []*po.Avatar, err error) {
	avatars = make([]*po.Avatar, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&avatars).Error
	return
}

func (r *avatarRepository) TxSaveAvatar(tx *gorm.DB, avatar *po.Avatar) (err error) {
	err = tx.Save(avatar).Error
	return
}
