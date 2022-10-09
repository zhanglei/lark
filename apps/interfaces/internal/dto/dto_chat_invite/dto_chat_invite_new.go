package dto_chat_invite

type NewChatInviteReq struct {
	ChatId   int64 `json:"chat_id" validate:"omitempty,gte=0"`        // chat ID
	ChatType int32 `json:"chat_type" validate:"required,gte=1,lte=2"` // 1:私聊/2:群聊
	//InitiatorUid  int64  `json:"initiator_uid" validate:"required,gt=0"`    // 发起人 UID
	InviteeUid    int64  `json:"invitee_uid" validate:"required,gt=0"` // 被邀请人 UID
	InvitationMsg string `json:"invitation_msg" validate:"required"`   // 邀请消息
}
