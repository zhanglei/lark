package ctrl_chat

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_chat"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *ChatCtrl) NewGroupChat(ctx *gin.Context) {
	var (
		params = new(dto_chat.NewGroupChatReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatService.NewGroupChat(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
