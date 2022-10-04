package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat_invite/internal/config"
	"lark/apps/chat_invite/internal/server"
	"lark/apps/chat_invite/internal/server/chat_invite"
	"lark/apps/chat_invite/internal/service"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(chat_invite.NewChatInviteServer)
	container.Provide(service.NewChatInviteService)
	container.Provide(repo.NewChatInviteRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
