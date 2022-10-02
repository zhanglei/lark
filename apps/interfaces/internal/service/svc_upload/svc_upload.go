package svc_upload

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_upload"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xminio"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"mime/multipart"
)

/*
https://www.cnblogs.com/peteremperor/p/16301336.html
https://www.cnblogs.com/liuqingzheng/p/16124105.html

https://www.lanol.cn/post/599.html
*/
type UploadService interface {
	UploadPhoto(ctx *gin.Context, req *dto_upload.UploadPhotoReq) (resp *xhttp.Resp)
	Presigned(ctx *gin.Context, req *dto_upload.PresignedReq) (resp *xhttp.Resp)
}

type uploadService struct {
	cfg *config.Config
}

func NewUploadService(cfg *config.Config) UploadService {
	return &uploadService{cfg: cfg}
}

func (s *uploadService) UploadPhoto(ctx *gin.Context, req *dto_upload.UploadPhotoReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		fileHeader *multipart.FileHeader
		file       multipart.File
		err        error
	)
	fileHeader, err = ctx.FormFile(constant.UPLOAD_PART_NAME_PHOTO)
	if err != nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED, err.Error())
		return
	}
	file, err = fileHeader.Open()
	if err != nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_OPEN_UPLOAD_FILE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_OPEN_UPLOAD_FILE_FAILED, err.Error())
		return
	}
	defer func() {
		file.Close()
	}()

	switch req.PhotoType {
	case constant.PHOTO_TYPE_AVATAR:
		var (
			new        *utils.NewPhoto
			resultList *xminio.PutResultList
		)
		new = utils.CropAvatar(file, s.cfg.Minio.PhotoDirectory, utils.NewUUID())
		if new.Error != nil {
			return
		}
		resultList = xminio.FPutPhotoListToMinio(new.List)
		if resultList.Err != nil {
			return
		}
	}
	return
}
