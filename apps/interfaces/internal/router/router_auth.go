package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"
	"lark/apps/interfaces/internal/service/svc_auth"
)

func registerAuthRouter(group *gin.RouterGroup) {
	var svc svc_auth.AuthService
	dig.Invoke(func(s svc_auth.AuthService) {
		svc = s
	})
	ctrl := ctrl_auth.NewAuthCtrl(svc)
	router := group.Group("auth")
	router.POST("login", ctrl.Login)
	router.POST("register", ctrl.Register)
}
