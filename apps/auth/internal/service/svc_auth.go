package service

import (
	"context"
	"lark/apps/auth/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_auth"
)

type AuthService interface {
	Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error)
	Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error)
}

type authService struct {
	cfg        *config.Config
	authRepo   repo.AuthRepository
	avatarRepo repo.AvatarRepository
}

func NewAuthService(cfg *config.Config, authRepo repo.AuthRepository, avatarRepo repo.AvatarRepository) AuthService {
	return &authService{cfg: cfg, authRepo: authRepo, avatarRepo: avatarRepo}
}
