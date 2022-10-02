package dig

import (
	"go.uber.org/dig"
	"lark/apps/interfaces/internal/config"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	provideAuth()
	provideUser()
	provideChat()
	provideUpload()
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
