package service

import (
	"context"
	"lark/domain/repo"
	"lark/pkg/proto/pb_req"
)

type ChatInviteService interface {
	NewChatRequest(ctx context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, err error)
	ChatRequestHandler(ctx context.Context, req *pb_req.ChatRequestHandlerReq) (resp *pb_req.ChatRequestHandlerResp, err error)
	ChatRequestList(ctx context.Context, req *pb_req.ChatRequestListReq) (resp *pb_req.ChatRequestListResp, err error)
}

type chatInviteService struct {
	chatInviteRepo repo.ChatInviteRepository
}

func NewChatInviteService(chatInviteRepo repo.ChatInviteRepository) ChatInviteService {
	return &chatInviteService{chatInviteRepo: chatInviteRepo}
}
