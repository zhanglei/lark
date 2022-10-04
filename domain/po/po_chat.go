package po

import "lark/pkg/entity"

type Chat struct {
	entity.GormEntityTs
	ChatId string `gorm:"column:chat_id;primary_key" json:"chat_id"` // chat ID
}
