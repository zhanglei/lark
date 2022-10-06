package dto_auth

import (
	"lark/apps/interfaces/internal/dto/dto_user"
	"lark/pkg/proto/pb_enum"
)

type LoginReq struct {
	AccountType      int32                 `json:"account_type" validate:"gte=1,lte=2"`       // 登录类型 1:手机号 2:lark账户
	Platform         pb_enum.PLATFORM_TYPE `json:"platform" validate:"gte=1,lte=2"`           // 平台 1:iOS 2:安卓
	Account          string                `json:"account" validate:"required,min=1,max=20"`  // 手机号/lark账户
	Udid             string                `json:"udid" validate:"required"`                  // 设备唯一编号
	VerificationCode string                `json:"verification_code"`                         // 验证码
	Password         string                `json:"password" validate:"required,min=8,max=20"` // 密码
}

type AuthResp struct {
	Token    *TokenInfo         `json:"token"`
	UserInfo *dto_user.UserInfo `json:"user_info"`
	Server   *ServerInfo        `json:"server"`
}

type TokenInfo struct {
	Token  string `json:"token"`  // 用户token
	Expire int64  `json:"expire"` // token过期时间戳（秒）
}

type ServerInfo struct {
	ServerId int32  `json:"server_id"` // 服务器ID
	Address  string `json:"address"`   // 服务器地址
}
