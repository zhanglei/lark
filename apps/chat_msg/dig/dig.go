package dig

import (
	"go.uber.org/dig"
	"lark/apps/chat_msg/internal/config"
	"lark/apps/chat_msg/internal/server"
	"lark/apps/chat_msg/internal/server/chat_msg"
	"lark/apps/chat_msg/internal/service"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(chat_msg.NewChatMessageServer)
	container.Provide(service.NewChatMessageService)
	container.Provide(repo.NewChatMessageRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
