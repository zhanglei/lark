package dos

type ChatMember struct {
	Uid      int64  `json:"uid"`      // 用户ID 系统生成
	LarkId   string `json:"lark_id"`  // 账户ID 用户设置
	Nickname string `json:"nickname"` // 昵称
}
