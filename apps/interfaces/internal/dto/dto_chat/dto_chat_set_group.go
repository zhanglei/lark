package dto_chat

import "lark/apps/interfaces/internal/dto/dto_kv"

type SetGroupChatReq struct {
	ChatId int64             `json:"chat_id" validate:"omitempty,gte=0"` // chat ID
	Kvs    *dto_kv.KeyValues `json:"kvs" validate:"required"`            // 更新字段
}
