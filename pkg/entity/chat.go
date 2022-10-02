package entity

type Chat struct {
	GormEntityTs
	ChatId string `gorm:"column:chat_id;primary_key" json:"chat_id"` // chat ID
}
