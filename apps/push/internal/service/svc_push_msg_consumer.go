package service

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"lark/apps/push/internal/logic"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"strconv"
)

var (
	isTest      bool
	testHashmap = make(map[string]string)
)

func init() {
	isTest = false
	if len(testHashmap) == 0 {
		for i := 1; i <= 10000; i++ {
			//member.Uid, member.Platform, member.ServerId, member.Mute
			testHashmap[strconv.Itoa(i)] = fmt.Sprintf("%d,%d,%d,%d", i, 1, 1, 0)
		}
	}
}

func (s *pushService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_NEW:
		err = s.pushMessage(msg)
		return
	case constant.CONST_MSG_KEY_RECALL:
		return
	default:
		return
	}
}

func (s *pushService) pushMessage(buf []byte) (err error) {
	var (
		inbox = new(pb_mq.InboxMessage)
	)
	if err = proto.Unmarshal(buf, inbox); err != nil {
		xlog.Warn(ERROR_CODE_PUSH_PROTOCOL_UNMARSHAL_ERR, ERROR_PUSH_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	s.chatMessagePush(inbox)
	return
}

func (s *pushService) chatMessagePush(inbox *pb_mq.InboxMessage) {
	//TODO:测试
	if isTest == true {
		s.groupChatMessagePushTest(inbox)
		return
	}

	var (
		key           string
		hashmap       map[string]string
		serverMembers map[int64][]*pb_obj.Int64Array
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(inbox.Msg.ChatId)
	hashmap = xredis.HGetAll(key)
	if len(hashmap) > 0 {
		serverMembers = logic.GetOnlinePushMembersFromHash(hashmap)
	} else {
		serverMembers = logic.GetOnlinePushMembersFromList(s.GetPushMemberList(inbox.Msg.ChatId))
	}
	s.onlinePushMessage(inbox, serverMembers)
}

func (s *pushService) groupChatMessagePushTest(inbox *pb_mq.InboxMessage) {
	var (
		serverMembers map[int64][]*pb_obj.Int64Array
	)
	serverMembers = logic.GetOnlinePushMembersFromHash(testHashmap)
	s.onlinePushMessage(inbox, serverMembers)
}

func (s *pushService) GetPushMemberList(chatId int64) (list []string) {
	var (
		userListReq = &pb_chat_member.GetPushMemberListReq{ChatId: chatId}
		resp        *pb_chat_member.GetPushMemberListResp
	)
	resp = s.chatMemberClient.GetPushMemberList(userListReq)
	if resp == nil {
		xlog.Warn(ERROR_CODE_PUSH_GRPC_SERVICE_FAILURE, ERROR_PUSH_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	return resp.List
}

func (s *pushService) onlinePushMessage(inbox *pb_mq.InboxMessage, serverMembers map[int64][]*pb_obj.Int64Array) {
	if len(serverMembers) == 0 {
		return
	}
	var (
		serverId int64
		req      *pb_gw.OnlinePushMessageReq
		pushResp *pb_gw.OnlinePushMessageResp
	)
	//TODO:serverId分组推送
	for serverId, _ = range serverMembers {
		req = &pb_gw.OnlinePushMessageReq{
			Topic:    inbox.Topic,
			SubTopic: inbox.SubTopic,
			Members:  serverMembers[serverId],
			Msg:      inbox.Msg,
		}
		pushResp = s.messageGatewayClient.OnlinePushMessage(req)
		if pushResp == nil {
			xlog.Warn(ERROR_CODE_PUSH_GRPC_SERVICE_FAILURE, ERROR_PUSH_GRPC_SERVICE_FAILURE)
		}
	}
	return
}
