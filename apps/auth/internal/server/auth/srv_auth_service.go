package auth

import (
	"context"
	"lark/pkg/proto/pb_auth"
)

func (s *authServer) Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error) {
	return s.authService.Register(ctx, req)
}

func (s *authServer) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error) {
	return s.authService.Login(ctx, req)
}
