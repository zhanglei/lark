package ctrl_auth

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (ctrl *AuthCtrl) Login(ctx *gin.Context) {
	var (
		params = new(dto_auth.LoginReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = ctx.BindJSON(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED, err.Error())
		return
	}
	if err = utils.Struct(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, err.Error())
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.authService.Login(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
