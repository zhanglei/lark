package svc_chat_invite

import (
	chat_invite_client "lark/apps/chat_invite/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_chat_invite"
	"lark/pkg/xhttp"
)

type ChatInviteService interface {
	NewChatInviteReq(params *dto_chat_invite.NewChatInviteReq, uid int64) (resp *xhttp.Resp)
	ChatInviteList(params *dto_chat_invite.ChatInviteListReq) (resp *xhttp.Resp)
	ChatInviteHandle(params *dto_chat_invite.ChatInviteHandleReq, uid int64) (resp *xhttp.Resp)
}

type chatInviteService struct {
	chatInviteClient chat_invite_client.ChatInviteClient
}

func NewChatInviteService() ChatInviteService {
	conf := config.GetConfig()
	chatInviteClient := chat_invite_client.NewChatInviteClient(conf.Etcd, conf.ChatInviteServer, conf.Jaeger, conf.Name)
	return &chatInviteService{chatInviteClient: chatInviteClient}
}
