package constant

//无状态
const (
	RK_LOGIN_TYPE            = "LOGIN_TYPE:"
	RK_MSG_CLI_MSG_ID        = "MSG:CLI_MSG_ID:" // 缓存客户端消息ID + ChatId + CliMsgId
	RK_MSG_SRV_MSG_ID        = "MSG:SRV_MSG_ID:" // 缓存服务端消息ID
	RK_MSG_SENDER_ID         = "MSG:SENDER_ID:"
	RK_MSG_CACHE             = "MSG:CACHE:" // 消息缓存 + ChatId + seqId
	RK_MSG_MAX_SEQ_ID        = "MSG:MAX_SEQ_ID:"
	RK_USER_LOGIN_PLATFORM   = "USER:LOGIN:PLATFORM:"   // 用户登录平台
	RK_USER_CHAT_MEMBER_INFO = "USER:CHAT_MEMBER_INFO:" // 聊天中需要用到的用户信息
)

//有状态
const (
	RK_SYNC_JWT                         = "JWT:"
	RK_SYNC_CHAT_MEMBERS_UID_LIST       = "CHAT:MEMBERS_UID_LIST:"       // 缓存群成员uid列表(表:chat_members.uid) + ChatId
	RK_SYNC_CHAT_MEMBERS_PUSH_CONF_HASH = "CHAT:MEMBERS_PUSH_CONF_HASH:" // 缓存chat成员推送配置(表:chat_members) + ChatId + UID
	RK_SYNC_CHAT_MEMBERS_SETTINGS_HASH  = "CHAT:MEMBERS_SETTINGS_HASH:"
	RK_SYNC_CHAT_MEMBERS_INFO_HASH      = "CHAT:MEMBERS_INFO_HASH:" // chat成员信息(表:chat_members) + ChatId + UID
	RK_SYNC_USERS_INFO                  = "USER:INFO:"              // 用户信息(表:users+user_avatars) + UID
	RK_SYNC_USERS_ONLINE_HASH           = "USER:ONLINE:"            // 在线用户数量 server_id + 数量
)
