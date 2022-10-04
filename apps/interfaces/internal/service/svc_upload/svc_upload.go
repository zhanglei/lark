package svc_upload

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_upload"
	user_client "lark/apps/user/client"
	"lark/pkg/common/xgopool"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xminio"
	"lark/pkg/common/xresize"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_user"
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
	UploadPhoto(ctx *gin.Context, req *dto_upload.UploadPhotoReq, uid int64) (resp *xhttp.Resp)
	Presigned(ctx *gin.Context, req *dto_upload.PresignedReq) (resp *xhttp.Resp)
}

type uploadService struct {
	cfg        *config.Config
	userClient user_client.UserClient
}

func NewUploadService(conf *config.Config) UploadService {
	userClient := user_client.NewUserClient(conf.Etcd, conf.UserServer, conf.Jaeger, conf.Name)
	return &uploadService{userClient: userClient, cfg: conf}
}

func (s *uploadService) UploadPhoto(ctx *gin.Context, req *dto_upload.UploadPhotoReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		fileHeader *multipart.FileHeader
		file       multipart.File
		reply      *pb_user.SetUserAvatarResp
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
			photos     *xresize.Photos
			resultList *xminio.PutResultList
			pr         *xminio.PutResult
			avatarReq  = &pb_user.SetUserAvatarReq{Uid: uid}
		)
		photos = xresize.CropAvatar(file, s.cfg.Minio.PhotoDirectory)
		if photos.Error != nil {
			resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_CROP_PHOTO_FAILED, xhttp.ERROR_HTTP_CROP_PHOTO_FAILED)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_CROP_PHOTO_FAILED, xhttp.ERROR_HTTP_CROP_PHOTO_FAILED, photos.Error.Error())
			return
		}
		resultList = xminio.FPutPhotoListToMinio(photos)
		if resultList.Err != nil {
			resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED, resultList.Err.Error())
			return
		}
		for _, pr = range resultList.List {
			switch photos.Maps[pr.Info.Key].Tag {
			case xresize.PhotoTagSmall:
				avatarReq.AvatarSmall = pr.Info.Key
			case xresize.PhotoTagMedium:
				avatarReq.AvatarMedium = pr.Info.Key
			case xresize.PhotoTagLarge:
				avatarReq.AvatarLarge = pr.Info.Key
			}
			path := photos.Maps[pr.Info.Key].Path
			xgopool.Go(func() {
				utils.Remove(path)
			})
		}
		reply = s.userClient.SetUserAvatar(avatarReq)
		if reply == nil {
			resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			return
		}
		if reply.Code > 0 {
			resp.SetRespInfo(reply.Code, reply.Msg)
			xlog.Warn(reply.Code, reply.Msg)
			return
		}
		resp.Data = reply.Avatar
	}
	return
}
