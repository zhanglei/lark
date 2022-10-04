package dig

import (
	"lark/apps/interfaces/internal/service/svc_chat_member"
)

func provideChatMember() {
	//container.Provide(ctrl_chat_member.NewChatMemberCtrl)
	container.Provide(svc_chat_member.NewChatMemberService)
}
