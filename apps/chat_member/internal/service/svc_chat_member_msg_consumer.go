package service

import (
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xredis"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/utils"
	"math"
)

func (s *chatMemberService) MessageHandler(msg []byte, msgKey string) (err error) {
	var (
		online = new(pb_mq.UserOnline)
		u      = entity.NewMysqlUpdate()
		w      = entity.NewMysqlWhere()
		list   []*pb_chat_member.ChatMemberPushConfig
	)
	proto.Unmarshal(msg, online)
	u.Query += " AND uid = ?"
	u.Args = append(u.Args, online.Uid)

	u.Set("server_id", online.ServerId)
	u.Set("platform", online.Platform)
	err = s.chatMemberUserRepo.UpdateChatMember(u)
	if err != nil {
		return
	}

	w.Query += " AND uid = ?"
	w.Args = append(w.Args, online.Uid)
	list, err = s.chatMemberUserRepo.ChatMemberPushConfigList(w)
	if err != nil {
		return
	}

	err = s.cacheMemberPushConfig(list)
	return
}

func (s *chatMemberService) cacheMemberPushConfig(list []*pb_chat_member.ChatMemberPushConfig) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		step          = 200
		consumerCount = int(math.Ceil(float64(len(list)) / float64(step)))
		errChan       = make(chan error, consumerCount)
		i             int
		j             int
	)

	for i = 0; i < consumerCount; i++ {
		var (
			minIndex = i * step
			maxIndex = (i + 1) * step
		)
		if maxIndex > len(list) {
			maxIndex = len(list)
		}
		go func(min, max int, configs []*pb_chat_member.ChatMemberPushConfig, errCh chan error) {
			var (
				cfg     *pb_chat_member.ChatMemberPushConfig
				jsonStr string
				key     string
				er      error
			)
			defer func() {
				errChan <- er
			}()
			for j = min; j < max; j++ {
				cfg = configs[j]
				jsonStr, er = utils.Marshal(cfg)
				if er != nil {
					break
				}
				er = xredis.HSetNX(key, utils.Int64ToStr(cfg.ChatId), jsonStr)
				if er != nil {
					break
				}
			}
		}(minIndex, maxIndex, list, errChan)
	}
	for i = 0; i < consumerCount; i++ {
		err = <-errChan
		if err != nil {
			break
		}
	}
	return
}
