package dto_upload

type UploadPhotoReq struct {
	PhotoType int32 `form:"photo_type" json:"photo_type"`
}

type UploadPhotoResp struct {
	Small  string `json:"small"`  // 小图
	Medium string `json:"medium"` // 中图
	Large  string `json:"large"`  // 大图
	Origin string `json:"origin"` // 原始图
}

type PresignedReq struct {
	FileType string `form:"file_type" json:"file_type"`
}

type PresignedResp struct {
	Url string `form:"url" json:"url"`
}
