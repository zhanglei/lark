package dig

import (
	"go.uber.org/dig"
	"lark/apps/wallet/internal/config"
	"lark/apps/wallet/internal/server"
	"lark/apps/wallet/internal/server/wallet"
	"lark/apps/wallet/internal/service"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(wallet.NewWalletServer)
	container.Provide(service.NewWalletService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
