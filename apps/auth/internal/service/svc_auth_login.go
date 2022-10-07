package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_avatar"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func setLoginResp(resp *pb_auth.LoginResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *authService) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, _ error) {
	resp = &pb_auth.LoginResp{UserInfo: &pb_user.UserInfo{Avatar: &pb_avatar.AvatarInfo{}}, Token: new(pb_auth.Token)}
	var (
		w      = entity.NewMysqlWhere()
		user   *po.User
		avatar *po.Avatar
		err    error
	)
	switch req.AccountType {
	case pb_enum.ACCOUNT_TYPE_MOBILE:
		w.Query += " AND mobile = ?"
		w.Args = append(w.Args, req.Account)
	case pb_enum.ACCOUNT_TYPE_LARK:
		w.Query += " AND lark_id = ?"
		w.Args = append(w.Args, req.Account)
	default:
		// 登录类型错误
		setLoginResp(resp, ERROR_CODE_AUTH_ACCOUNT_TYPE_ERR, ERROR_AUTH_ACCOUNT_TYPE_ERR)
		xlog.Warn(ERROR_CODE_AUTH_ACCOUNT_TYPE_ERR, ERROR_AUTH_ACCOUNT_TYPE_ERR)
		return
	}
	w.Query += " AND password = ?"
	w.Args = append(w.Args, utils.MD5(req.Password))
	user, err = s.authRepo.VerifyUserIdentity(w)
	if err != nil {
		setLoginResp(resp, ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR, ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR)
		xlog.Warn(ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR, ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR, err.Error())
		return
	}
	if user.Uid == 0 {
		setLoginResp(resp, ERROR_CODE_AUTH_USER_DOES_NOT_EXIST, ERROR_AUTH_USER_DOES_NOT_EXIST)
		xlog.Warn(ERROR_CODE_AUTH_USER_DOES_NOT_EXIST, ERROR_AUTH_USER_DOES_NOT_EXIST, req.String())
		return
	}

	w.Reset()
	w.SetFilter("owner_id=?", user.Uid)
	w.SetFilter("owner_type=?", int32(pb_enum.AVATAR_OWNER_USER_AVATAR))
	avatar, err = s.avatarRepo.Avatar(w)
	if err != nil {
		setLoginResp(resp, ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_QUERY_DB_FAILED, ERROR_AUTH_QUERY_DB_FAILED, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	copier.Copy(resp.UserInfo.Avatar, avatar)
	resp.Token.Token, resp.Token.Expire = xjwt.CreateToken(user.Uid, int32(req.Platform))
	return
}
