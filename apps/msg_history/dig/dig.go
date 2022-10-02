package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_history/internal/config"
	"lark/apps/msg_history/internal/domain/repo"
	"lark/apps/msg_history/internal/server"
	"lark/apps/msg_history/internal/server/msg_history"
	"lark/apps/msg_history/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(repo.NewMessageHistoryRepository)
	container.Provide(msg_history.NewMessageHistoryServer)
	container.Provide(service.NewMessageHistoryService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
