package utils

func PrivateChatSessionKey(chatId int64, uid1 int64, uid2 int64) string {
	var (
		sessionKey string
	)
	if uid1 > uid2 {
		sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid1) + "-" + Int64ToStr(uid2)
	} else {
		sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid2) + "-" + Int64ToStr(uid1)
	}
	return MD5(sessionKey)
}

func GroupChatSessionKey(chatId int64, uid int64) string {
	var (
		sessionKey string
	)
	sessionKey = Int64ToStr(chatId) + "-" + Int64ToStr(uid)
	return MD5(sessionKey)
}
