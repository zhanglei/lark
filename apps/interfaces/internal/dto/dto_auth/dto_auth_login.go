package dto_auth

import "lark/pkg/proto/pb_enum"

type LoginReq struct {
	AccountType      int32                 `json:"account_type" validate:"gte=1,lte=2"`       // 登录类型 1:手机号 2:lark账户
	Platform         pb_enum.PLATFORM_TYPE `json:"platform" validate:"gte=1,lte=2"`           // 平台 1:iOS 2:安卓
	Account          string                `json:"account" validate:"required,min=1,max=20"`  // 手机号/lark账户
	Udid             string                `json:"udid" validate:"required"`                  // 设备唯一编号
	VerificationCode string                `json:"verification_code"`                         // 验证码
	Password         string                `json:"password" validate:"required,min=8,max=20"` // 密码
}

type LoginResp struct {
	Token    *TokenInfo  `json:"token"`
	UserInfo *UserInfo   `json:"user_info"`
	Server   *ServerInfo `json:"server"`
}

type TokenInfo struct {
	Token  string `json:"token"`  // 用户token
	Expire int64  `json:"expire"` // token过期时间戳（秒）
}

type UserInfo struct {
	Uid       int64  `json:"uid"`        // uid
	LarkId    string `json:"lark_id"`    // 账户ID
	Status    int32  `json:"status"`     // 用户状态
	Nickname  string `json:"nickname"`   // 昵称
	Firstname string `json:"firstname"`  // firstname
	Lastname  string `json:"lastname"`   // lastname
	Gender    int32  `json:"gender"`     // 性别
	BirthTs   int64  `json:"birth_ts"`   // 生日
	Email     string `json:"email"`      // Email
	Mobile    string `json:"mobile"`     // 手机号
	AvatarKey string `json:"avatar_key"` // 头像
	CityId    int64  `json:"city_id"`    // 城市ID
}

type ServerInfo struct {
	ServerId int32  `json:"server_id"` // 服务器ID
	Address  string `json:"address"`   // 服务器地址
}
