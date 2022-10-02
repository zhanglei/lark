package dig

import (
	"go.uber.org/dig"
	"lark/apps/msg_hot/internal/config"
	"lark/apps/msg_hot/internal/domain/repo"
	"lark/apps/msg_hot/internal/server"
	"lark/apps/msg_hot/internal/server/msg_hot"
	"lark/apps/msg_hot/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(repo.NewMessageHotRepository)
	container.Provide(msg_hot.NewMessageHotServer)
	container.Provide(service.NewMessageHotService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
