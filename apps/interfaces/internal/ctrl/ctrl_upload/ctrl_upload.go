package ctrl_upload

import (
	"lark/apps/interfaces/internal/service/svc_upload"
)

type UploadCtrl struct {
	svc svc_upload.UploadService
}

func NewUploadCtrl(svc svc_upload.UploadService) *UploadCtrl {
	return &UploadCtrl{svc: svc}
}
