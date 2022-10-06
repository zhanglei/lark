package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_chat"
	"lark/apps/interfaces/internal/service/svc_chat"
)

func registerChatRouter(group *gin.RouterGroup) {
	var svc svc_chat.ChatService
	dig.Invoke(func(s svc_chat.ChatService) {
		svc = s
	})
	ctrl := ctrl_chat.NewChatCtrl(svc)
	router := group.Group("chat")
	router.POST("new", ctrl.NewGroupChat)
	router.POST("set", ctrl.SetGroupChat)
}
