package repo

import (
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type UserAvatarRepository interface {
	UserAvatar(w *entity.MysqlWhere) (avatar *po.UserAvatar, err error)
	UserAvatarList(w *entity.MysqlWhere) (avatars []*po.UserAvatar, err error)
	TxSaveAvatar(tx *gorm.DB, avatar *po.UserAvatar) (err error)
}

type userAvatarRepository struct {
}

func NewUserAvatarRepository() UserAvatarRepository {
	return &userAvatarRepository{}
}

func (r *userAvatarRepository) UserAvatar(w *entity.MysqlWhere) (avatar *po.UserAvatar, err error) {
	avatar = new(po.UserAvatar)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(avatar).Error
	return
}

func (r *userAvatarRepository) UserAvatarList(w *entity.MysqlWhere) (avatars []*po.UserAvatar, err error) {
	avatars = make([]*po.UserAvatar, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&avatars).Error
	return
}

func (r *userAvatarRepository) TxSaveAvatar(tx *gorm.DB, avatar *po.UserAvatar) (err error) {
	err = tx.Save(avatar).Error
	return
}
