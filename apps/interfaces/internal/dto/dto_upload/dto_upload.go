package dto_upload

type UploadAvatarReq struct {
	OwnerId   int64 `form:"owner_id" json:"owner_id"`
	OwnerType int32 `form:"owner_type" json:"owner_type"`
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

type ObjectStorage struct {
	Bucket      string `json:"bucket"`
	Key         string `json:"key"`
	ETag        string `json:"e_tag"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
	//Format      string `json:"format"`
	//UUID        string `json:"uuid"`
	FileName string `json:"file_name"`
	Tag      string `json:"tag"`
}
