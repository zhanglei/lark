package service

import (
	"context"
	"lark/apps/chat/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_chat"
)

type ChatService interface {
	NewChat(ctx context.Context, req *pb_chat.NewChatReq) (resp *pb_chat.NewChatResp, err error)
}

type chatService struct {
	cfg      *config.Config
	chatRepo repo.ChatRepository
}

func NewChatService(cfg *config.Config, chatRepo repo.ChatRepository) ChatService {
	svc := &chatService{cfg: cfg, chatRepo: chatRepo}
	return svc
}
