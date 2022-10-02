package dig

import (
	"lark/apps/interfaces/internal/ctrl/ctrl_chat_msg"
	"lark/apps/interfaces/internal/service/svc_chat_msg"
)

func provideChat() {
	container.Provide(ctrl_chat_msg.NewChatMessageCtrl)
	container.Provide(svc_chat_msg.NewChatMessageService)
}
