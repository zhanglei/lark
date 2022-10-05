package dto_chat_invite

type ChatInviteListReq struct {
	Uid          int64 `json:"uid" validate:"required,gt=0"`
	Role         int32 `json:"role" validate:"required,gt=0"` // 角色
	MaxInviteId  int64 `json:"max_invite_id" validate:"required,gt=0"`
	HandleResult int32 `json:"handle_result" validate:"required,gte=1,lte=2"` // 结果
	Limit        int32 `json:"limit" validate:"required,gt=0"`
}
