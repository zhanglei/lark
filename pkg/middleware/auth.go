package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"lark/pkg/common/xjwt"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			err   error
			token *jwt.Token
			ok    bool
		)
		token, err = xjwt.ParseFromCookie(ctx)
		if err != nil {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_JWT_TOKEN_ERR, err.Error())
			return
		}
		claims := jwt.MapClaims{}
		for key, value := range token.Claims.(jwt.MapClaims) {
			claims[key] = value
		}
		if _, ok = claims[constant.USER_UID]; ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
			return
		}
		if _, ok = claims[constant.USER_PLATFORM]; ok == false || utils.TryToInt(claims[constant.USER_PLATFORM]) == 0 {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
			return
		}
		ctx.Set(constant.USER_UID, claims[constant.USER_UID])
		ctx.Set(constant.USER_PLATFORM, claims[constant.USER_PLATFORM])
	}
}
