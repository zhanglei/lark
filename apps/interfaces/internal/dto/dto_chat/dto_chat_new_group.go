package dto_chat

type NewGroupChatReq struct {
	Title   string  `json:"title" validate:"required"`     // 标题
	About   string  `json:"about" validate:"required"`     // About
	UidList []int64 `json:"uid_list" validate:"omitempty"` // 邀请人员uid列表
}
