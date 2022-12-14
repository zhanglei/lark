package service

const (
	ERROR_CODE_USER_ACCOUNT_TYPE_ERR        int32 = 21001
	ERROR_CODE_USER_ACCOUNT_OR_PASSWORD_ERR int32 = 21002
	ERROR_CODE_USER_QUERY_DB_FAILED         int32 = 21003
	ERROR_CODE_USER_REDIS_GET_FAILED        int32 = 21004
	ERROR_CODE_USER_SET_AVATAR_FAILED       int32 = 21005
)
const (
	ERROR_USER_ACCOUNT_TYPE_ERR        = "登录类型错误"
	ERROR_USER_ACCOUNT_OR_PASSWORD_ERR = "账户或密码错误"
	ERROR_USER_QUERY_DB_FAILED         = "查询失败"
	ERROR_USER_REDIS_GET_FAILED        = "读取redis缓存失败"
	ERROR_USER_SET_AVATAR_FAILED       = "设置用户头像失败"
)

const (
	ERROR_CODE_USER_REGISTER_ERR = 21101
)
const (
	ERROR_USER_REGISTER_TYPE_ERR = "注册失败"
)
