package ctrl_upload

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_upload"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *UploadCtrl) Presigned(ctx *gin.Context) {
	var (
		params dto_upload.PresignedReq
		resp   *xhttp.Resp
		err    error
	)
	if err = ctx.ShouldBindQuery(&params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED, err.Error())
		return
	}
	resp = ctrl.svc.Presigned(ctx, &params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
