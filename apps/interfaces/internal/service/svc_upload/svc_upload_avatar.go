package svc_upload

import (
	"github.com/gin-gonic/gin"
	avatar_client "lark/apps/avatar/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_upload"
	"lark/pkg/common/xgopool"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xminio"
	"lark/pkg/common/xresize"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_avatar"
	"lark/pkg/proto/pb_enum"
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
	UploadAvatar(ctx *gin.Context, req *dto_upload.UploadAvatarReq, uid int64) (resp *xhttp.Resp)
	Presigned(ctx *gin.Context, req *dto_upload.PresignedReq) (resp *xhttp.Resp)
}

type uploadService struct {
	cfg          *config.Config
	avatarClient avatar_client.AvatarClient
}

func NewUploadService(conf *config.Config) UploadService {
	avatarClient := avatar_client.NewAvatarClient(conf.Etcd, conf.AvatarServer, conf.Jaeger, conf.Name)
	return &uploadService{avatarClient: avatarClient, cfg: conf}
}

func (s *uploadService) UploadAvatar(ctx *gin.Context, req *dto_upload.UploadAvatarReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		fileHeader *multipart.FileHeader
		file       multipart.File
		reply      *pb_avatar.SetAvatarResp
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

	var (
		photos     *xresize.Photos
		resultList *xminio.PutResultList
		pr         *xminio.PutResult
		avatarReq  = &pb_avatar.SetAvatarReq{
			OwnerId:   req.OwnerId,
			OwnerType: pb_enum.AVATAR_OWNER(req.OwnerType),
		}
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
	reply = s.avatarClient.SetAvatar(avatarReq)
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
	return
}
