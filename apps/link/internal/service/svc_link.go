package service

import (
	"context"
	chat_member_client "lark/apps/chat_member/client"
	user_client "lark/apps/user/client"
	"lark/pkg/proto/pb_link"
)

type LinkService interface {
	UserOnline(ctx context.Context, req *pb_link.UserOnlineReq) (resp *pb_link.UserOnlineResp, err error)
	UserOffline(ctx context.Context, req *pb_link.UserOfflineReq) (resp *pb_link.UserOfflineResp, err error)
}

type linkService struct {
	userClient       user_client.UserClient
	chatMemberClient chat_member_client.ChatMemberClient
}

func NewLinkService() LinkService {
	return &linkService{}
}
