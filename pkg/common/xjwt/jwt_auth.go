package xjwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"lark/pkg/constant"
	"time"
)

func CreateToken(uid int64, platform int32) (tokenString string, expireIn int64) {
	var (
		token  *jwt.Token
		claims jwt.MapClaims
		err    error
	)
	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)

	expireIn = constant.JWT_TOKEN_EXPIRE_IN
	claims["iss"] = constant.JWT_ISSUER
	claims["exp"] = time.Now().Add(time.Duration(expireIn) * time.Second).Unix()
	claims["orig_iat"] = time.Now().Unix()
	claims[constant.USER_UID] = uid
	claims[constant.USER_PLATFORM] = platform
	tokenString, err = token.SignedString([]byte(constant.JWT_TOKEN_KEY))
	if err != nil {
		expireIn = -1
		return
	}
	tokenString = constant.JWT_FIELD + tokenString
	return
}

// CreateRefreshToken 创建刷新token
func CreateRefreshToken(uid int64, platform int32) (tokenString string, expireIn int64) {
	var (
		token  *jwt.Token
		claims jwt.MapClaims
		err    error
	)
	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)
	expireIn = constant.JWT_REFRESH_TOKEN_EXPIRE_IN
	claims["iss"] = constant.JWT_ISSUER
	claims["exp"] = time.Now().Add(time.Duration(expireIn) * time.Second).Unix()
	claims["orig_iat"] = time.Now().Unix()
	claims[constant.USER_UID] = uid
	claims[constant.USER_PLATFORM] = platform
	tokenString, err = token.SignedString([]byte(constant.JWT_REFRESH_TOKEN_KEY))
	if err != nil {
		expireIn = -1
	}
	return
}

// ParseRefreshToken 解析刷新Token
func ParseRefreshToken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_REFRESH_TOKEN_KEY), nil
	})
	return
}

func ParseFromHeader(ctx *gin.Context) (res *jwt.Token, err error) {
	res, err = request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.JWT_TOKEN_KEY), nil
		})
	if err == request.ErrNoTokenInRequest {
		token := ctx.Query("token")
		res, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(constant.JWT_TOKEN_KEY), nil
		})
	}
	return
}

func ParseFromCookie(ctx *gin.Context) (*jwt.Token, error) {
	token := ctx.Query("token")
	cookie, _ := ctx.Cookie("jwt")
	tokenStr := cookie
	if token != "" {
		tokenStr = token
	}
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_TOKEN_KEY), nil
	})
}

func ParseFromToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_TOKEN_KEY), nil
	})
}
