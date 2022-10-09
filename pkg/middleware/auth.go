package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			err      error
			token    *jwt.Token
			ok       bool
			uid      interface{}
			platform interface{}
			uuid     interface{}
			uuidVal  string
			uuidKey  string
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
		if uid, ok = claims[constant.USER_UID]; ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
			return
		}
		if platform, ok = claims[constant.USER_PLATFORM]; ok == false || utils.TryToInt(claims[constant.USER_PLATFORM]) == 0 {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, xhttp.ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
			return
		}
		if uuid, ok = claims[constant.USER_JWT_UUID]; ok == false {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_JWT_TOKEN_UUID_DOESNOT_EXIST, xhttp.ERROR_HTTP_JWT_TOKEN_UUID_DOESNOT_EXIST)
			return
		}
		uuidKey = constant.RK_SYNC_JWT_UUID + utils.ToString(uid)
		if uuidVal, err = xredis.Get(uuidKey); err != nil {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_TOKEN_EXPIRED, xhttp.ERROR_HTTP_TOKEN_EXPIRED)
			return
		}
		if uuidVal != utils.ToString(uuid) {
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_TOKEN_EXPIRED, xhttp.ERROR_HTTP_TOKEN_EXPIRED)
			return
		}
		ctx.Set(constant.USER_UID, uid)
		ctx.Set(constant.USER_PLATFORM, platform)
	}
}
