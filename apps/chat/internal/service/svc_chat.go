package service

import (
	"context"
	"lark/apps/chat/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_chat"
)

type ChatService interface {
	NewGroupChat(ctx context.Context, req *pb_chat.NewGroupChatReq) (resp *pb_chat.NewGroupChatResp, err error)
	SetGroupChat(ctx context.Context, req *pb_chat.SetGroupChatReq) (resp *pb_chat.SetGroupChatResp, err error)
}

type chatService struct {
	cfg            *config.Config
	chatRepo       repo.ChatRepository
	chatInviteRepo repo.ChatInviteRepository
	chatMemberRepo repo.ChatMemberRepository
	userRepo       repo.UserRepository
}

func NewChatService(cfg *config.Config,
	chatRepo repo.ChatRepository,
	chatInviteRepo repo.ChatInviteRepository,
	chatMemberRepo repo.ChatMemberRepository,
	userRepo repo.UserRepository) ChatService {
	svc := &chatService{cfg: cfg,
		chatRepo:       chatRepo,
		chatInviteRepo: chatInviteRepo,
		chatMemberRepo: chatMemberRepo,
		userRepo:       userRepo}
	return svc
}
