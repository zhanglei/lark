package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_user"
)

func setRegisterResp(resp *pb_auth.RegisterResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *authService) Register(_ context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, _ error) {
	resp = &pb_auth.RegisterResp{UserInfo: new(pb_user.UserInfo), Token: new(pb_auth.Token)}
	var (
		user = new(entity.User)
		err  error
	)
	copier.Copy(user, req)
	err = s.authRepo.Create(user)
	if err != nil {
		setRegisterResp(resp, ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR)
		xlog.Warn(ERROR_CODE_AUTH_REGISTER_ERR, ERROR_AUTH_REGISTER_TYPE_ERR, err.Error())
		return
	}
	copier.Copy(resp.UserInfo, user)
	resp.Token.Token, resp.Token.Expire = xjwt.CreateToken(user.Uid, int32(req.RegPlatform))
	return
}
