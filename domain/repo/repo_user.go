package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
)

type UserRepository interface {
	Create(user *po.User) (err error)
	VerifyUserIdentity(w *entity.MysqlWhere) (user *po.User, err error)
	UserInfo(w *entity.MysqlWhere) (user *po.User, err error)
	UserList(w *entity.MysqlWhere) (list []*po.User, err error)
	UpdateUser(u *entity.MysqlUpdate) (err error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

/*
存:传指针对象，Create时不需要&，同时会Out表中的数据
读:返回指针对象，Find时不需要&
需要不为nil
*/

func (r *userRepository) Create(user *po.User) (err error) {
	user.Uid = xsnowflake.NewSnowflakeID()
	if user.LarkId == "" {
		user.LarkId = xsnowflake.DefaultLarkId()
	}
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *userRepository) VerifyUserIdentity(w *entity.MysqlWhere) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(user).Error
	return
}

func (r *userRepository) UserList(w *entity.MysqlWhere) (list []*po.User, err error) {
	list = make([]*po.User, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&list).Error
	return
}

func (r *userRepository) UserInfo(w *entity.MysqlWhere) (user *po.User, err error) {
	user = new(po.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(&user).Error
	return
}

func (r *userRepository) UpdateUser(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	err = db.Model(po.User{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}
