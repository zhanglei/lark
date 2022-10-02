package request

import (
	"context"
	"lark/pkg/proto/pb_req"
)

func (s *requestServer) NewChatRequest(ctx context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, err error) {
	return s.requestService.NewChatRequest(ctx, req)
}
