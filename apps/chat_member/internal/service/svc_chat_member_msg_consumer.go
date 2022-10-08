package service

import (
	"google.golang.org/protobuf/proto"
	"lark/domain/do"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
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
		list   []*do.ChatMemberInfo
	)
	proto.Unmarshal(msg, online)
	if online.Uid == 0 {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_MISS_USER_INFO, ERROR_CHAT_MEMBER_MISS_USER_INFO)
		return
	}
	u.SetFilter("uid = ?", online.Uid)
	u.Set("server_id", online.ServerId)
	u.Set("platform", online.Platform)
	err = s.chatMemberRepo.UpdateChatMember(u)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_UPDATE_VALUE_FAILED, ERROR_CHAT_MEMBER_UPDATE_VALUE_FAILED, err.Error())
		return
	}
	w.SetFilter("uid = ?", online.Uid)
	list, err = s.chatMemberRepo.ChatMemberList(w)
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED, ERROR_CHAT_MEMBER_QUERY_DB_FAILED, err.Error())
		return
	}
	err = s.cachePushMembers(list)
	return
}

func (s *chatMemberService) cachePushMembers(list []*do.ChatMemberInfo) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		step          = 1000
		consumerCount = int(math.Ceil(float64(len(list)) / float64(step)))
		errChan       = make(chan error, consumerCount)
		i             int
		j             int
	)
	// TODO:携程优化
	for i = 0; i < consumerCount; i++ {
		var (
			minIndex = i * step
			maxIndex = (i + 1) * step
		)
		if maxIndex > len(list) {
			maxIndex = len(list)
		}
		go func(min, max int, members []*do.ChatMemberInfo, errCh chan error) {
			var (
				m       *do.ChatMemberInfo
				member  *pb_chat_member.PushMember
				jsonStr string
				key     string
				er      error
			)
			defer func() {
				errChan <- er
			}()
			for j = min; j < max; j++ {
				m = members[j]
				member = &pb_chat_member.PushMember{
					Uid:      m.Uid,
					ServerId: m.ServerId,
					Platform: m.Platform,
					Mute:     m.Mute,
				}
				jsonStr, er = utils.Marshal(member)
				if er != nil {
					break
				}
				key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_CONF_HASH + utils.Int64ToStr(m.ChatId)
				er = xredis.HSetNX(key, utils.Int64ToStr(member.Uid), jsonStr)
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
