package dig

import (
	"lark/apps/interfaces/internal/service/svc_auth"
)

func provideAuth() {
	//container.Provide(ctrl_auth.NewAuthCtrl)
	container.Provide(svc_auth.NewAuthService)
}
