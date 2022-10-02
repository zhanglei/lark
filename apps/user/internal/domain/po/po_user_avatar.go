package po

import "lark/pkg/entity"

type UserAvatar struct {
	entity.GormEntityTs
	Uid          int64  `gorm:"column:uid;primary_key" json:"uid"`         // 用户ID 系统生成
	AvatarSmall  string `gorm:"column:avatar_small" json:"avatar_small"`   // 小图72
	AvatarMedium string `gorm:"column:avatar_medium" json:"avatar_medium"` // 中图240
	AvatarLarge  string `gorm:"column:avatar_large" json:"avatar_large"`   // 大图640
	AvatarOrigin string `gorm:"column:avatar_origin" json:"avatar_origin"` // 原始图
}
