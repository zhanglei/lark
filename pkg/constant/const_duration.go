package constant

import "time"

const (
	CONST_DURATION_LOGIN_PLATFORM_SECOND = time.Duration(7*24*60*60) * time.Second
)

const (
	CONST_DURATION_REDIS_TIMEOUT = 500 * time.Millisecond
)
const (
	CONST_DURATION_USER_INFO_SECOND = time.Duration(1*24*60*60) * time.Second //用户信息缓存时间
)
const (
	CONST_DURATION_CHAT_GROUP_UID_LIST_SECOND    = time.Duration(1*24*60*60) * time.Second //群uid列表缓存时间
	CONST_DURATION_CHAT_USER_SETTING_SECOND      = time.Duration(1*24*60*60) * time.Second
	CONST_DURATION_CHAT_MEMBER_BASIC_LIST_SECOND = time.Duration(1*24*60*60) * time.Second

	//CONST_DURATION_CHAT_MEMBER_SESSION_SECOND = time.Duration(7*24*60*60) * time.Second //聊天成员信息缓存时间
	CONST_DURATION_CHAT_MEMBER_INFO_SECOND = time.Duration(7*24*60*60) * time.Second //聊天成员信息缓存时间
)

const (
	CONST_DURATION_MSG_ID_SECOND    = time.Duration(12*60*60) * time.Second //消息ID缓存时间,用于判断是否是重复消息
	CONST_DURATION_MSG_CACHE_SECOND = time.Duration(24*60*60) * time.Second //单个消息缓存时间
)
