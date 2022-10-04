package dig

import (
	"lark/apps/interfaces/internal/service/svc_user"
)

func provideUser() {
	//container.Provide(ctrl_user.NewUserCtrl)
	container.Provide(svc_user.NewUserService)
}
