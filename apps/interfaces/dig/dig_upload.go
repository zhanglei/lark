package dig

import (
	"lark/apps/interfaces/internal/service/svc_upload"
)

func provideUpload() {
	container.Provide(svc_upload.NewUploadService)
}
