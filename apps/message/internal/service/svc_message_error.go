package service

import "errors"

const (
	ERROR_CODE_MESSAGE_UNMARSHAL_ERR          int32 = 71001
	ERROR_CODE_MESSAGE_VALIDATOR_ERR          int32 = 71002
	ERROR_CODE_MESSAGE_REPEATED_MESSAGE       int32 = 71003
	ERROR_CODE_MESSAGE_ENQUEUE_FAILED         int32 = 71004
	ERROR_CODE_MESSAGE_RECORD_MSGID_FAILED    int32 = 71005
	ERROR_CODE_MESSAGE_GET_SEQ_ID_FAILED      int32 = 71006
	ERROR_CODE_MESSAGE_VERIFY_IDENTITY_FAILED int32 = 71007
	ERROR_CODE_MESSAGE_REDIS_GET_FAILED       int32 = 71008
	ERROR_CODE_MESSAGE_GRPC_SERVICE_FAILURE   int32 = 71009 // 服务故障
)

const (
	ERROR_MESSAGE_UNMARSHAL_ERR          = "消息反序列化错误"
	ERROR_MESSAGE_VALIDATOR_ERR          = "参数校验错误"
	ERROR_MESSAGE_REPEATED_MESSAGE       = "重复的消息"
	ERROR_MESSAGE_ENQUEUE_FAILED         = "消息入队失败"
	ERROR_MESSAGE_RECORD_MSGID_FAILED    = "缓存消息ID失败"
	ERROR_MESSAGE_GET_SEQ_ID_FAILED      = "生成 Sequence ID 失败"
	ERROR_MESSAGE_VERIFY_IDENTITY_FAILED = "验证身份失败"
	ERROR_MESSAGE_REDIS_GET_FAILED       = "读取redis缓存失败"
	ERROR_MESSAGE_GRPC_SERVICE_FAILURE   = "服务故障"
)

var (
	ERROR_MESSAGE_BODY_TEXT_EMPTY_ERR = errors.New("消息为空")
)
