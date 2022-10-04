package repos

import (
	"lark/domain/pos"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

type UserAvatarRepository interface {
	UserAvatar(w *entity.MysqlWhere) (avatar *pos.UserAvatar, err error)
	SaveAvatar(avatar *pos.UserAvatar) (err error)
}

type userAvatarRepository struct {
}

func NewUserAvatarRepository() UserAvatarRepository {
	return &userAvatarRepository{}
}

func (r *userAvatarRepository) UserAvatar(w *entity.MysqlWhere) (avatar *pos.UserAvatar, err error) {
	avatar = new(pos.UserAvatar)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(avatar).Error
	return
}

func (r *userAvatarRepository) SaveAvatar(avatar *pos.UserAvatar) (err error) {
	db := xmysql.GetDB()
	err = db.Save(avatar).Error
	return
}
