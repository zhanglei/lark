package dig

import (
	"go.uber.org/dig"
	"lark/apps/offline_push/internal/config"
	"lark/apps/offline_push/internal/server"
	"lark/apps/offline_push/internal/server/offline_push"
	"lark/apps/offline_push/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(offline_push.NewOfflinePushServer)
	container.Provide(service.NewOfflinePushService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
