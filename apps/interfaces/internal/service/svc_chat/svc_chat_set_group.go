package svc_chat

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/xhttp"
)

func (s *chatService) SetGroupChat(params *dto_chat.SetGroupChatReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat.SetGroupChatReq)
		reply   *pb_chat.SetGroupChatResp
	)
	copier.Copy(reqArgs, params)
	reply = s.chatClient.SetGroupChat(reqArgs)
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
