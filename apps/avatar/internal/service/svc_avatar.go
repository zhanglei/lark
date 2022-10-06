package service

import (
	"context"
	"lark/apps/avatar/internal/config"
	"lark/domain/repo"
	"lark/pkg/proto/pb_avatar"
)

type AvatarService interface {
	SetAvatar(ctx context.Context, req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp, err error)
}

type avatarService struct {
	cfg            *config.Config
	avatarRepo     repo.AvatarRepository
	chatMemberRepo repo.ChatMemberRepository
	chatRepo       repo.ChatRepository
}

func NewAvatarService(cfg *config.Config, avatarRepo repo.AvatarRepository,
	chatMemberRepo repo.ChatMemberRepository,
	chatRepo repo.ChatRepository) AvatarService {
	svc := &avatarService{cfg: cfg,
		avatarRepo:     avatarRepo,
		chatMemberRepo: chatMemberRepo,
		chatRepo:       chatRepo}
	return svc
}
