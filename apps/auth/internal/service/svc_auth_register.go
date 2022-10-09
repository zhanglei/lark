package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"lark/domain/po"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xmysql"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_avatar"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func setRegisterResp(resp *pb_auth.RegisterResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *authService) Register(_ context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, _ error) {
	resp = &pb_auth.RegisterResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_avatar.AvatarInfo{}}, Token: new(pb_auth.Token)}
	var (
		user   = new(po.User)
		avatar *po.Avatar
		tx     *gorm.DB
		err    error
	)
	copier.Copy(user, req)
	user.AvatarKey = constant.CONST_AVATAR_KEY_SMALL
	user.Password = utils.MD5(user.Password)
	user.ServerId = 1
	user.Platform = int(req.RegPlatform)

	tx = xmysql.GetTX()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	err = s.authRepo.TxCreate(tx, user)
	if err != nil {
		setRegisterResp(resp, ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR)
		xlog.Warn(ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR, err.Error())
		return
	}
	avatar = &po.Avatar{
		OwnerId:      user.Uid,
		OwnerType:    int(pb_enum.AVATAR_OWNER_USER_AVATAR),
		AvatarSmall:  constant.CONST_AVATAR_KEY_SMALL,
		AvatarMedium: constant.CONST_AVATAR_KEY_MEDIUM,
		AvatarLarge:  constant.CONST_AVATAR_KEY_LARGE,
	}
	err = s.avatarRepo.TxCreate(tx, avatar)
	if err != nil {
		setRegisterResp(resp, ERROR_CODE_AUTH_INSERT_VALUE_FAILED, ERROR_AUTH_INSERT_VALUE_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_INSERT_VALUE_FAILED, ERROR_AUTH_INSERT_VALUE_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	copier.Copy(resp.UserInfo.Avatar, avatar)
	resp.Token.Token, resp.Token.Expire = xjwt.CreateToken(user.Uid, int32(req.RegPlatform))
	return
}
