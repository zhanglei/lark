package router

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	//engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	publicGroup := engine.Group("open")
	registerPublicRoutes(publicGroup)

	privateGroup := engine.Group("api")
	registerPrivateRouter(privateGroup)
}

func registerPublicRoutes(group *gin.RouterGroup) {
	registerAuthRouter(group)
}

func registerPrivateRouter(group *gin.RouterGroup) {
	group.Use(middleware.JwtAuth())
	registerUserRouter(group)
	registerChatMessageRouter(group)
	registerUploadRouter(group)
	registerChatMemberRouter(group)
	registerChatInviteRouter(group)
	registerChatRouter(group)
}
