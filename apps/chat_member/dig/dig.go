package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat_member/internal/config"
	"lark/apps/chat_member/internal/domain/repo"
	"lark/apps/chat_member/internal/server"
	"lark/apps/chat_member/internal/server/chat_member"
	"lark/apps/chat_member/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(chat_member.NewChatMemberServer)
	container.Provide(service.NewChatMemberService)
	container.Provide(repo.NewChatMemberRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
