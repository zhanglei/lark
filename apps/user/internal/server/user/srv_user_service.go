package user

import (
	"context"
	"lark/pkg/proto/pb_user"
)

func (s *userServer) GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	return s.userService.GetUserInfo(ctx, req)
}

func (s *userServer) GetChatUserInfo(ctx context.Context, req *pb_user.GetChatUserInfoReq) (resp *pb_user.GetChatUserInfoResp, err error) {
	return s.userService.GetChatUserInfo(ctx, req)
}

func (s *userServer) UserOnline(ctx context.Context, req *pb_user.UserOnlineReq) (resp *pb_user.UserOnlineResp, err error) {
	return s.userService.UserOnline(ctx, req)
}
