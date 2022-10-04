package pos

import "lark/pkg/entity"

type UserAvatar struct {
	entity.GormEntityTs
	Uid          int64  `gorm:"column:uid;primary_key" json:"uid"`         // 用户ID
	AvatarSmall  string `gorm:"column:avatar_small" json:"avatar_small"`   // 小图 72*72
	AvatarMedium string `gorm:"column:avatar_medium" json:"avatar_medium"` // 中图 240*240
	AvatarLarge  string `gorm:"column:avatar_large" json:"avatar_large"`   // 大图 640*640
}
