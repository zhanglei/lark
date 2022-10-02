package constant

const (
	JWT_ISSUER                  = "tKh3fMwWepCds9jbTuJ9IGYJYbPIZnj"
	JWT_TOKEN_KEY               = "lark_jwt_token_2022"
	JWT_TOKEN_EXPIRE_IN         = 3600 * 24 * 30
	JWT_REFRESH_TOKEN_EXPIRE_IN = 3600 * 24 * 60 // 刷新token的时长
	JWT_REFRESH_TOKEN_KEY       = "lark_jwt_refresh_token_2022"
	JWT_FIELD                   = "jwt="
)
