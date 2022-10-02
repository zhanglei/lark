package service

import (
	"context"
	"lark/apps/request/internal/domain/repo"
	"lark/pkg/proto/pb_req"
)

type RequestService interface {
	NewChatRequest(ctx context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, err error)
	ChatRequestHandler(ctx context.Context, req *pb_req.ChatRequestHandlerReq) (resp *pb_req.ChatRequestHandlerResp, err error)
	ChatRequestList(ctx context.Context, req *pb_req.ChatRequestListReq) (resp *pb_req.ChatRequestListResp, err error)
}

type requestService struct {
	requestRepo repo.RequestRepository
}

func NewRequestService(requestRepo repo.RequestRepository) RequestService {
	return &requestService{requestRepo: requestRepo}
}
