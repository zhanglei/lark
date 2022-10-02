package dig

import (
	"go.uber.org/dig"
	"lark/apps/push/internal/config"
	"lark/apps/push/internal/server"
	"lark/apps/push/internal/server/push"
	"lark/apps/push/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(push.NewPushServer)
	container.Provide(service.NewPushService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
