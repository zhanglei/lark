package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_user"
	"lark/apps/interfaces/internal/service/svc_user"
)

func registerUserRouter(group *gin.RouterGroup) {
	var svc svc_user.UserService
	dig.Invoke(func(s svc_user.UserService) {
		svc = s
	})
	ctrl := ctrl_user.NewUserCtrl(svc)
	router := group.Group("user")
	router.POST("list", ctrl.UserList)
}
