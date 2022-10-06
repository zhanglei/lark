package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_chat_msg"
	"lark/apps/interfaces/internal/service/svc_chat_msg"
)

func registerChatMessageRouter(group *gin.RouterGroup) {
	var svc svc_chat_msg.ChatMessageService
	dig.Invoke(func(s svc_chat_msg.ChatMessageService) {
		svc = s
	})
	ctrl := ctrl_chat_msg.NewChatMessageCtrl(svc)
	router := group.Group("chat_msg")
	router.GET("list", ctrl.GetChatMessages)
}
