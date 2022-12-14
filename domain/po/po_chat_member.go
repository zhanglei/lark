package po

import "lark/pkg/entity"

type ChatMember struct {
	entity.GormEntityTs
	ChatId          int64  `gorm:"column:chat_id;primary_key" json:"chat_id"`                  // chat ID
	ChatType        int    `gorm:"column:chat_type;NOT NULL" json:"chat_type"`                 // chat type 1:私聊/2:群聊
	Uid             int64  `gorm:"column:uid;NOT NULL" json:"uid"`                             // 用户ID
	RoleId          int    `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"`           // 角色ID
	Mute            int    `gorm:"column:mute;default:0;NOT NULL" json:"mute"`                 // 是否开启免打扰
	DisplayName     string `gorm:"column:display_name;NOT NULL" json:"display_name"`           // 显示名称
	MemberAvatarKey string `gorm:"column:member_avatar_key;NOT NULL" json:"member_avatar_key"` // member头像 72*72
	ChatAvatarKey   string `gorm:"column:chat_avatar_key;NOT NULL" json:"chat_avatar_key"`     // chat头像 72*72
	Sync            int    `gorm:"column:sync;default:0;NOT NULL" json:"sync"`                 // 是否同步用户信息
	Status          int    `gorm:"column:status;default:0;NOT NULL" json:"status"`             // chat状态
	Platform        int    `gorm:"column:platform;NOT NULL" json:"platform"`                   // 1:iOS 2:安卓
	ServerId        int    `gorm:"column:server_id;default:0;NOT NULL" json:"server_id"`       // 服务器ID
	Settings        string `gorm:"column:settings;NOT NULL" json:"settings"`                   // 用户设置
}
