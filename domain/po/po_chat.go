package po

import "lark/pkg/entity"

type Chat struct {
	entity.GormEntityTs
	ChatId    int64  `gorm:"column:chat_id;primary_key" json:"chat_id"`   // chat ID
	ChatHash  string `gorm:"column:chat_hash" json:"chat_hash"`           // chat hash值
	ChatType  int    `gorm:"column:chat_type;default:0" json:"chat_type"` // chat type 1:私聊/2:群聊
	AvatarKey string `gorm:"column:avatar_key" json:"avatar_key"`         // 小图 72*62
	Title     string `gorm:"column:title" json:"title"`                   // chat标题
	About     string `gorm:"column:about" json:"about"`                   // 关于
}
