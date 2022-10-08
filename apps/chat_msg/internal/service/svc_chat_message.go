package service

import (
	"context"
	"lark/apps/chat_msg/internal/config"
	"lark/domain/mrepo"
	"lark/domain/repo"
	"lark/pkg/proto/pb_chat_msg"
)

type ChatMessageService interface {
	GetChatMessages(_ context.Context, req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp, _ error)
}

type chatMessageService struct {
	conf            *config.Config
	chatMessageRepo repo.ChatMessageRepository
	messageHotRepo  mrepo.MessageHotRepository
}

func NewChatMessageService(conf *config.Config, chatMessageRepo repo.ChatMessageRepository, messageHotRepo mrepo.MessageHotRepository) ChatMessageService {
	return &chatMessageService{conf: conf, chatMessageRepo: chatMessageRepo, messageHotRepo: messageHotRepo}
}
