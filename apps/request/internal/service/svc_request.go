package service

import (
	"context"
	"lark/domain/repos"
	"lark/pkg/proto/pb_req"
)

type RequestService interface {
	NewChatRequest(ctx context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, err error)
	ChatRequestHandler(ctx context.Context, req *pb_req.ChatRequestHandlerReq) (resp *pb_req.ChatRequestHandlerResp, err error)
	ChatRequestList(ctx context.Context, req *pb_req.ChatRequestListReq) (resp *pb_req.ChatRequestListResp, err error)
}

type requestService struct {
	requestRepo repos.RequestRepository
}

func NewRequestService(requestRepo repos.RequestRepository) RequestService {
	return &requestService{requestRepo: requestRepo}
}
