package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_mq"
)

func (s *userService) MessageHandler(msg []byte, msgKey string) (err error) {
	var (
		online = new(pb_mq.UserOnline)
		u      = entity.NewMysqlUpdate()
	)
	proto.Unmarshal(msg, online)
	u.Query += " AND uid = ?"
	u.Args = append(u.Args, online.Uid)

	u.Set("server_id", online.ServerId)
	u.Set("platform", online.Platform)
	err = s.userRepo.UpdateUser(u)
	return
}
