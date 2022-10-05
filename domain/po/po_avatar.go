package po

import "lark/pkg/entity"

type Avatar struct {
	entity.GormEntityTs
	AvatarId     int64  `gorm:"column:avatar_id;primary_key" json:"avatar_id"`
	OwnerId      int64  `gorm:"column:owner_id;NOT NULL" json:"owner_id"`      // 用户ID/ChatID
	OwnerType    int    `gorm:"column:owner_type;default:0" json:"owner_type"` // 1:用户头像 2:群头像
	AvatarSmall  string `gorm:"column:avatar_small" json:"avatar_small"`       // 小图 72*62
	AvatarMedium string `gorm:"column:avatar_medium" json:"avatar_medium"`     // 中图 240*240
	AvatarLarge  string `gorm:"column:avatar_large" json:"avatar_large"`       // 大图 640*640
}
