package dos

type ChatMemberPushConfig struct {
	ChatId   int64 `json:"chat_id"`   // chat ID
	Uid      int64 `json:"uid"`       // 用户ID
	Mute     int   `json:"mute"`      // 是否开启免打扰
	Platform int   `json:"platform"`  // 1:iOS 2:安卓
	ServerId int   `json:"server_id"` // 服务器ID
}
