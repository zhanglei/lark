package service

import (
	"context"
	"lark/domain/repo"
	"lark/pkg/proto/pb_invite"
)

type ChatInviteService interface {
	NewChatInvite(ctx context.Context, req *pb_invite.NewChatInviteReq) (resp *pb_invite.NewChatInviteResp, err error)
	ChatInviteHandle(ctx context.Context, req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp, err error)
	ChatInviteList(ctx context.Context, req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp, err error)
}

type chatInviteService struct {
	chatInviteRepo repo.ChatInviteRepository
	userRepo       repo.UserRepository
	avatarRepo     repo.AvatarRepository
	chatMemberRepo repo.ChatMemberRepository
	chatRepo       repo.ChatRepository
}

func NewChatInviteService(
	chatInviteRepo repo.ChatInviteRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	chatRepo repo.ChatRepository) ChatInviteService {
	return &chatInviteService{
		chatInviteRepo: chatInviteRepo,
		userRepo:       userRepo,
		avatarRepo:     avatarRepo,
		chatMemberRepo: chatMemberRepo,
		chatRepo:       chatRepo}
}
