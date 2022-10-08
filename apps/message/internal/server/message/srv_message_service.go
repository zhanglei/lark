package message

import (
	"context"
	"lark/pkg/proto/pb_msg"
)

func (s *messageServer) SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error) {
	return s.messageService.SendChatMessage(ctx, req)
}

func (s *messageServer) MessageOperation(ctx context.Context, req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp, err error) {
	return s.messageService.MessageOperation(ctx, req)
}
