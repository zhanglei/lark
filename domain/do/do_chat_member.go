package do

import "lark/pkg/proto/pb_enum"

type ChatMemberInfo struct {
	ChatId   int64                 `json:"chat_id"`   // chat ID
	Uid      int64                 `json:"uid"`       // 用户ID
	Mute     pb_enum.MUTE_TYPE     `json:"mute"`      // 是否开启免打扰
	Platform pb_enum.PLATFORM_TYPE `json:"platform"`  // 1:iOS 2:安卓
	ServerId int32                 `json:"server_id"` // 服务器ID
}
