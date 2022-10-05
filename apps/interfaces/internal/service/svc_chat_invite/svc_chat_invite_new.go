package svc_chat_invite

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_invite"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_invite"
	"lark/pkg/xhttp"
)

func (s *chatInviteService) NewChatInviteReq(params *dto_chat_invite.NewChatInviteReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_invite.NewChatInviteReq)
		reply   *pb_invite.NewChatInviteResp
	)
	copier.Copy(reqArgs, params)
	reply = s.chatInviteClient.NewChatInvite(reqArgs)
	if reply == nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetRespInfo(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	return
}
