package dig

import (
	"go.uber.org/dig"
	"lark/apps/request/internal/config"
	"lark/apps/request/internal/domain/repo"
	"lark/apps/request/internal/server"
	"lark/apps/request/internal/server/request"
	"lark/apps/request/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(request.NewRequestServer)
	container.Provide(service.NewRequestService)
	container.Provide(repo.NewRequestRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}