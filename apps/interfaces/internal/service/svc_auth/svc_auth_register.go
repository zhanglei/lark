package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/xhttp"
)

func (s *authService) Register(params *dto_auth.RegisterReq) (resp *xhttp.Resp) {
	var (
		req          = new(pb_auth.RegisterReq)
		reply        *pb_auth.RegisterResp
		registerResp = new(dto_auth.RegisterResp)
	)
	reply = s.authClient.Register(req)
	if reply == nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetRespInfo(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
	}
	copier.Copy(registerResp, reply)
	resp.Data = registerResp
	return
}
