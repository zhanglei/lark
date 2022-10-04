package service

import (
	"context"
	"lark/domain/repo"
	"lark/pkg/proto/pb_invite"
)

type ChatInviteService interface {
	NewChatInvite(_ context.Context, req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp, err error)
	ChatInviteHandle(ctx context.Context, req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp, err error)
	ChatInviteList(_ context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, err error)
}

type chatInviteService struct {
	chatInviteRepo repo.ChatInviteRepository
}

func NewChatInviteService(chatInviteRepo repo.ChatInviteRepository) ChatInviteService {
	return &chatInviteService{chatInviteRepo: chatInviteRepo}
}
