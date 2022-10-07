package xgin

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func GetUid(ctx *gin.Context) (uid int64) {
	var (
		value  any
		exists bool
	)
	value, exists = ctx.Get(constant.USER_UID)
	if exists == false {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	uid, _ = utils.ToInt64(value)
	if uid == 0 {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	return
}
