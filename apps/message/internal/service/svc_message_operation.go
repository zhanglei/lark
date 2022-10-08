package service

import (
	"context"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
)

func (s *messageService) MessageOperation(ctx context.Context, req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp, err error) {
	resp = new(pb_msg.MessageOperationResp)
	switch req.Operation {
	case pb_enum.MSG_OPERATION_RECALL:
	case pb_enum.MSG_OPERATION_URGENT:
	default:
		return
	}
	return
}
