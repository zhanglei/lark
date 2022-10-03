package svc_upload

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_upload"
	"lark/pkg/common/xgopool"
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
			photos     *utils.Photos
			resultList *xminio.PutResultList
			pr         *xminio.PutResult
			list       []*dto_upload.ObjectStorage
		)
		photos = utils.CropAvatar(file, s.cfg.Minio.PhotoDirectory)
		if photos.Error != nil {
			return
		}
		resultList = xminio.FPutPhotoListToMinio(photos)
		if resultList.Err != nil {
			return
		}
		list = make([]*dto_upload.ObjectStorage, 3)
		for _, pr = range resultList.List {
			os := &dto_upload.ObjectStorage{
				Bucket:      pr.Info.Bucket,
				Key:         pr.Info.Key,
				ETag:        pr.Info.ETag,
				Size:        pr.Info.Size,
				ContentType: photos.ContentType,
				FileName:    fileHeader.Filename,
			}
			os.Tag = photos.Maps[pr.Info.Key].Tag
			switch os.Tag {
			case "small":
				list[0] = os
			case "medium":
				list[1] = os
			case "large":
				list[2] = os
			}
			path := photos.Maps[pr.Info.Key].Path
			xgopool.Go(func() {
				utils.Remove(path)
			})
		}
		// TODO:入库
		resp.Data = list
	}
	return
}
