package link

import (
	"context"
	"lark/pkg/proto/pb_link"
)

func (s *linkServer) UserOnline(ctx context.Context, req *pb_link.UserOnlineReq) (resp *pb_link.UserOnlineResp, err error) {
	return s.linkService.UserOnline(ctx, req)
}

func (s *linkServer) UserOffline(ctx context.Context, req *pb_link.UserOfflineReq) (resp *pb_link.UserOfflineResp, err error) {
	return s.linkService.UserOffline(ctx, req)
}
