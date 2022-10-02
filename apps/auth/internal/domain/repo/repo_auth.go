package repo

import (
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/entity"
)

type AuthRepository interface {
	Create(user *entity.User) (err error)
	VerifyUserIdentity(w *entity.MysqlWhere) (user *entity.User, err error)
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

/*
存:传指针对象，Create时不需要&，同时会Out表中的数据
读:返回指针对象，Find时不需要&
需要不为nil
*/

func (r *authRepository) Create(user *entity.User) (err error) {
	user.Uid = xsnowflake.NewSnowflakeID()
	if user.LarkId == "" {
		user.LarkId = xsnowflake.DefaultLarkId()
	}
	db := xmysql.GetDB()
	err = db.Create(user).Error
	return
}

func (r *authRepository) VerifyUserIdentity(w *entity.MysqlWhere) (user *entity.User, err error) {
	user = new(entity.User)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(user).Error
	return
}
