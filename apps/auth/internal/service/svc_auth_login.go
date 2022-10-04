package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
)

func setLoginResp(resp *pb_auth.LoginResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *authService) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, _ error) {
	resp = &pb_auth.LoginResp{UserInfo: new(pb_user.UserInfo), Token: new(pb_auth.Token)}
	var (
		w    = entity.NewMysqlWhere()
		user *po.User
		err  error
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
	w.Args = append(w.Args, req.Password)
	user, err = s.authRepo.VerifyUserIdentity(w)
	if err != nil {
		setLoginResp(resp, ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR, ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR)
		xlog.Warn(ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR, ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	resp.Token.Token, resp.Token.Expire = xjwt.CreateToken(user.Uid, int32(req.Platform))
	return
}
