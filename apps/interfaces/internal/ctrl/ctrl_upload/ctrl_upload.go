package ctrl_upload

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_upload"
	"lark/apps/interfaces/internal/service/svc_upload"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

type UploadCtrl struct {
	svc svc_upload.UploadService
}

func NewUploadCtrl(svc svc_upload.UploadService) *UploadCtrl {
	return &UploadCtrl{svc: svc}
}

func (ctrl *UploadCtrl) UploadPhoto(ctx *gin.Context) {
	var (
		params   dto_upload.UploadPhotoReq
		resp     *xhttp.Resp
		keyValue any
		exists   bool
		uid      int
		err      error
	)
	if err = ctx.Bind(&params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED, err.Error())
		return
	}
	keyValue, exists = ctx.Get(constant.USER_UID)
	if exists == false {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	uid, _ = utils.ToInt(keyValue)
	if uid == 0 {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	resp = ctrl.svc.UploadPhoto(ctx, &params, int64(uid))
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

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
