package svc_chat_msg

import (
	"lark/apps/chat_msg/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	"lark/pkg/xhttp"
)

type ChatMessageService interface {
	GetChatMessages(req *dto_chat_msg.GetChatMessagesReq) (resp *xhttp.Resp)
}

type chatMessageService struct {
	chatMessageClient chat_msg_client.ChatMessageClient
}

func NewChatMessageService() ChatMessageService {
	conf := config.GetConfig()
	chatMessageClient := chat_msg_client.NewChatMessageClient(conf.Etcd, conf.ChatMsgServer, conf.Jaeger, conf.Name)
	return &chatMessageService{chatMessageClient: chatMessageClient}
}
