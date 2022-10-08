package service

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"lark/apps/push/internal/logic"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_mq"
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
			testHashmap[strconv.Itoa(i)] = fmt.Sprintf("{\"chat_id\":%d,\"uid\":%d,\"mute\":0,\"platform\":1,\"server_id\":1}", 3333336666669999990, i)
		}
	}
}

func (s *pushService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_NEW:
		err = s.PushMessage(msg)
		return
	case constant.CONST_MSG_KEY_RECALL:
		return
	default:
		return
	}
}

func (s *pushService) PushMessage(buf []byte) (err error) {
	var (
		inbox         = new(pb_mq.InboxMessage)
		member        *pb_chat_member.PushMember
		serverMembers map[int32][]*pb_chat_member.PushMember
	)
	if err = proto.Unmarshal(buf, inbox); err != nil {
		xlog.Warn(ERROR_CODE_PUSH_PROTOCOL_UNMARSHAL_ERR, ERROR_PUSH_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	switch inbox.Msg.ChatType {
	case pb_enum.CHAT_TYPE_PRIVATE:
		member = s.GetPushMember(inbox.Msg.ChatId, inbox.Msg.ReceiverId)
		if member == nil {
			return
		}
		serverMembers = logic.GetOnlinePushMembersFromList([]*pb_chat_member.PushMember{member})
		s.onlinePushMessage(inbox, serverMembers)
	case pb_enum.CHAT_TYPE_GROUP:
		s.groupChatMessagePush(inbox)
	}
	return
}

func (s *pushService) GetPushMember(chatId int64, uid int64) (member *pb_chat_member.PushMember) {
	//TODO:测试
	if isTest {
		member = &pb_chat_member.PushMember{
			Uid:      uid,
			Mute:     pb_enum.MUTE_TYPE_CLOSED,
			Platform: pb_enum.PLATFORM_TYPE_IOS,
			ServerId: 1,
		}
		return
	}
	var (
		key  string
		list []interface{}
		req  *pb_chat_member.GetPushMemberReq
		resp *pb_chat_member.GetPushMemberResp
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(chatId)
	list = xredis.HMGet(key, utils.Int64ToStr(uid))
	if len(list) == 1 && list[0] != nil {
		member = new(pb_chat_member.PushMember)
		utils.Unmarshal(list[0].(string), member)
		return
	}
	req = &pb_chat_member.GetPushMemberReq{ChatId: chatId, Uid: uid}
	resp = s.chatMemberClient.GetPushMember(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_PUSH_GRPC_SERVICE_FAILURE, ERROR_PUSH_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	member = resp.Member
	return
}

func (s *pushService) groupChatMessagePush(inbox *pb_mq.InboxMessage) {
	//TODO:测试
	if isTest == true {
		s.groupChatMessagePushTest(inbox)
		return
	}

	var (
		key           string
		hashmap       map[string]string
		serverMembers map[int32][]*pb_chat_member.PushMember
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_MEMBER_HASH + utils.Int64ToStr(inbox.Msg.ChatId)
	hashmap = xredis.HGetAll(key)
	if len(hashmap) > 1 {
		serverMembers = logic.GetOnlinePushMembersFromHash(hashmap)
	} else {
		serverMembers = logic.GetOnlinePushMembersFromList(s.GetPushMemberList(inbox.Msg.ChatId))
	}
	s.onlinePushMessage(inbox, serverMembers)
}

func (s *pushService) groupChatMessagePushTest(inbox *pb_mq.InboxMessage) {
	var (
		serverMembers map[int32][]*pb_chat_member.PushMember
	)
	serverMembers = logic.GetOnlinePushMembersFromHash(testHashmap)
	s.onlinePushMessage(inbox, serverMembers)
}

func (s *pushService) GetPushMemberList(chatId int64) (settings []*pb_chat_member.PushMember) {
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

func (s *pushService) onlinePushMessage(inbox *pb_mq.InboxMessage, serverMembers map[int32][]*pb_chat_member.PushMember) {
	if len(serverMembers) == 0 {
		return
	}
	var (
		serverId int32
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
