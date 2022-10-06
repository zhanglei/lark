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
		uidVal int
		exists bool
	)
	value, exists = ctx.Get(constant.USER_UID)
	if exists == false {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	uidVal, _ = utils.ToInt(value)
	if uidVal == 0 {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
		return
	}
	uid = int64(uidVal)
	return
}
