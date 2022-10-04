package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/pos"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/proto/pb_req"
)

func setNewChatRequestResp(resp *pb_req.NewChatRequestResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *requestService) NewChatRequest(_ context.Context, req *pb_req.NewChatRequestReq) (resp *pb_req.NewChatRequestResp, _ error) {
	resp = new(pb_req.NewChatRequestResp)
	var (
		request = new(pos.ChatRequest)
		err     error
	)
	copier.Copy(request, req)
	request.RequestId = xsnowflake.NewSnowflakeID()
	err = s.requestRepo.RequestCreate(request)
	if err != nil {
		setNewChatRequestResp(resp, ERROR_CODE_REQUEST_INSERT_VALUE_FAILED, ERROR_REQUEST_INSERT_VALUE_FAILED)
		xlog.Warn(resp, ERROR_CODE_REQUEST_INSERT_VALUE_FAILED, ERROR_REQUEST_INSERT_VALUE_FAILED, err)
		return
	}
	return
}
