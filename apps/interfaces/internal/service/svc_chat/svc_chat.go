package svc_chat

import (
	chat_client "lark/apps/chat/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_chat"
	"lark/pkg/xhttp"
)

type ChatService interface {
	NewGroupChat(params *dto_chat.NewGroupChatReq, uid int64) (resp *xhttp.Resp)
	SetGroupChat(params *dto_chat.SetGroupChatReq) (resp *xhttp.Resp)
}

type chatService struct {
	chatClient chat_client.ChatClient
}

func NewChatService() ChatService {
	conf := config.GetConfig()
	chatClient := chat_client.NewChatClient(conf.Etcd, conf.ChatServer, conf.Jaeger, conf.Name)
	return &chatService{chatClient: chatClient}
}
