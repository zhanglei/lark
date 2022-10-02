package xhttp

const (
	ERROR_HTTP_REQ_DESERIALIZE_FAILED = "请求参数反序列化错误"
	ERROR_HTTP_REQ_PARAM_ERR          = "请求参数错误"
	//ErrorHttpReqNotAuthorized         = "没有授权"
	//ErrorHttpRegisterFailed           = "注册失败"
	ERROR_HTTP_USER_ID_DOESNOT_EXIST  = "用户ID信息缺失"
	ERROR_HTTP_PLATFORM_DOESNOT_EXIST = "平台信息缺失"
	//ErrorHttpGetUserFailed        = "获取用户信息失败"
	//ErrorHttpAddFriendFailed          = "添加好友失败"
	ERROR_HTTP_SERVICE_FAILURE         = "服务故障"
	ERROR_HTTP_MESSAGE_ENQUEUE_FAILED  = "消息入队失败"
	ERROR_HTTP_PRESIGNED_FAILED        = "上传文件预先签署失败"
	ERROR_HTTP_READ_UPLOAD_FILE_FAILED = "读取上传文件失败"
	ERROR_HTTP_OPEN_UPLOAD_FILE_FAILED = "打开上传文件失败"
)