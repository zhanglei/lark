package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/pos"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_user"
)

func setUserAvatarResp(resp *pb_user.SetUserAvatarResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *userService) SetUserAvatar(ctx context.Context, req *pb_user.SetUserAvatarReq) (resp *pb_user.SetUserAvatarResp, _ error) {
	resp = &pb_user.SetUserAvatarResp{Avatar: &pb_user.UserAvatar{}}
	var (
		avatar = new(pos.UserAvatar)
		err    error
	)
	copier.Copy(avatar, req)
	err = s.userAvatarRepo.SaveAvatar(avatar)
	if err != nil {
		setUserAvatarResp(resp, ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED)
		xlog.Warn(ERROR_CODE_USER_SET_AVATAR_FAILED, ERROR_USER_SET_AVATAR_FAILED, err.Error())
		return
	}
	copier.Copy(resp.Avatar, avatar)
	return
}
