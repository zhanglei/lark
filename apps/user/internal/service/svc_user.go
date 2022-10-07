package service

import (
	"context"
	"lark/apps/user/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_user"
)

type UserService interface {
	GetUserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error)
	GetUserList(ctx context.Context, req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp, err error)
	GetChatUserInfo(ctx context.Context, req *pb_user.GetChatUserInfoReq) (resp *pb_user.GetChatUserInfoResp, err error)
	UserOnline(ctx context.Context, req *pb_user.UserOnlineReq) (resp *pb_user.UserOnlineResp, err error)
}

type userService struct {
	cfg            *config.Config
	userRepo       repo.UserRepository
	avatarRepo     repo.AvatarRepository
	chatMemberRepo repo.ChatMemberRepository
}

func NewUserService(cfg *config.Config, userRepo repo.UserRepository, avatarRepo repo.AvatarRepository, chatMemberRepo repo.ChatMemberRepository) UserService {
	svc := &userService{cfg: cfg, userRepo: userRepo, avatarRepo: avatarRepo, chatMemberRepo: chatMemberRepo}
	return svc
}
