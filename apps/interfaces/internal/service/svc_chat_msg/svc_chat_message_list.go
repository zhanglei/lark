package svc_chat_msg

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/xhttp"
)

func (s *chatMessageService) GetChatMessages(req *dto_chat_msg.GetChatMessagesReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		getChatMessagesReq  = new(pb_chat_msg.GetChatMessagesReq)
		getChatMessagesResp *pb_chat_msg.GetChatMessagesResp
	)
	copier.Copy(getChatMessagesReq, req)
	getChatMessagesResp = s.chatMessageClient.GetChatMessages(getChatMessagesReq)
	if getChatMessagesResp == nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if getChatMessagesResp.Code > 0 {
		resp.SetRespInfo(getChatMessagesResp.Code, getChatMessagesResp.Msg)
		xlog.Warn(getChatMessagesResp.Code, getChatMessagesResp.Msg)
		return
	}
	resp.Data = getChatMessagesResp.List
	return
}
