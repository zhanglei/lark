package xgin

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func BindJSON(ctx *gin.Context, params interface{}) (err error) {
	if err = ctx.BindJSON(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		return
	}
	if err = utils.Struct(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	return
}

func ShouldBindQuery(ctx *gin.Context, params interface{}) (err error) {
	if err = ctx.ShouldBindQuery(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		return
	}
	if err = utils.Struct(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	return
}
