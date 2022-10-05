package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_upload"
	"lark/apps/interfaces/internal/service/svc_upload"
)

func registerUploadRouter(group *gin.RouterGroup) {
	var svc svc_upload.UploadService
	dig.Invoke(func(s svc_upload.UploadService) {
		svc = s
	})
	ctrl := ctrl_upload.NewUploadCtrl(svc)
	router := group.Group("upload")
	router.POST("avatar", ctrl.UploadAvatar)
	router.GET("presigned", ctrl.Presigned)
}
