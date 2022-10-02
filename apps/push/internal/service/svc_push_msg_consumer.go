package service

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"lark/apps/push/internal/logic"
	"lark/pkg/common/xgopool"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_push"
	"lark/pkg/utils"
	"strconv"
)

var (
	isTest      bool
	testHashmap = make(map[string]string)
)

func init() {
	isTest = true
	if len(testHashmap) == 0 {
		for i := 1; i <= 10000; i++ {
			testHashmap[strconv.Itoa(i)] = fmt.Sprintf("{\"chat_id\":%d,\"uid\":%d,\"mute\":0,\"platform\":1,\"server_id\":1}", 3333336666669999990, i)
		}
	}
}

func (s *pushService) MessageHandler(msg []byte, chatId string) (err error) {
	var (
		req     = new(pb_mq.InboxMessage)
		pushReq *pb_push.PushMessageReq
	)
	if err = proto.Unmarshal(msg, req); err != nil {
		xlog.Warn(ERROR_CODE_PUSH_PROTOCOL_UNMARSHAL_ERR, ERROR_PUSH_PROTOCOL_UNMARSHAL_ERR, err.Error())
		return
	}
	// 消息推送
	pushReq = &pb_push.PushMessageReq{
		Topic:    req.Topic,
		SubTopic: req.SubTopic,
		Msg:      req.Msg,
	}
	xgopool.Go(func() {
		s.PushMessage(pushReq)
	})
	return
}

func (s *pushService) PushMessage(req *pb_push.PushMessageReq) (err error) {
	var (
		conf *pb_chat_member.ChatMemberPushConfig
	)
	//conf = s.getChatMemberPushConfig(req.Msg.ChatId, req.Msg.SenderId)
	//if conf == nil {
	//	return
	//}
	//logic.PushMessageToMembers(req, []*pb_chat_member.ChatMemberPushConfig{conf}, s.privateChatMessagePush)
	switch req.Msg.ChatType {
	case pb_enum.CHAT_TYPE_PRIVATE:
		conf = s.getChatMemberPushConfig(req.Msg.ChatId, req.Msg.ReceiverId)
		if conf == nil {
			return
		}
		logic.PushMessageToMembers(req, []*pb_chat_member.ChatMemberPushConfig{conf}, s.chatMessagePush)
	case pb_enum.CHAT_TYPE_GROUP:
		s.groupChatMessagePush(req)
	}
	return
}

func (s *pushService) getChatMemberPushConfig(chatId int64, uid int64) (conf *pb_chat_member.ChatMemberPushConfig) {
	//TODO:测试
	if isTest {
		conf = &pb_chat_member.ChatMemberPushConfig{
			ChatId:   chatId,
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
		req  *pb_chat_member.GetChatMemberPushConfigReq
		resp *pb_chat_member.GetChatMemberPushConfigResp
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_CONF_HASH + utils.Int64ToStr(chatId)
	list = xredis.HMGet(key, utils.Int64ToStr(uid))
	if len(list) > 0 && list[0] != nil {
		conf = new(pb_chat_member.ChatMemberPushConfig)
		utils.Unmarshal(list[0].(string), conf)
		return
	}
	req = &pb_chat_member.GetChatMemberPushConfigReq{ChatId: chatId, Uid: uid}
	resp = s.chatMemberClient.GetChatMemberPushConfig(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_PUSH_GRPC_SERVICE_FAILURE, ERROR_PUSH_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
	conf = resp.Config
	return
}

func (s *pushService) groupChatMessagePush(req *pb_push.PushMessageReq) {
	//TODO:测试
	if isTest == true {
		s.groupChatMessagePushTest(req)
		return
	}

	var (
		key     string
		hashmap map[string]string
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_PUSH_CONF_HASH + utils.Int64ToStr(req.Msg.ChatId)
	hashmap = xredis.HGetAll(key)
	if len(hashmap) > 1 {
		logic.PushMessageToHashmapMembers(req, hashmap, s.chatMessagePush)
	} else {
		logic.PushMessageToMembers(req, s.getChatMemberPushConfigList(req.Msg.ChatId), s.chatMessagePush)
	}
}

func (s *pushService) groupChatMessagePushTest(req *pb_push.PushMessageReq) {
	logic.PushMessageToHashmapMembers(req, testHashmap, s.chatMessagePush)
}

func (s *pushService) getChatMemberPushConfigList(chatId int64) (settings []*pb_chat_member.ChatMemberPushConfig) {
	var (
		userListReq = &pb_chat_member.GetChatMemberPushConfigListReq{ChatId: chatId}
		resp        *pb_chat_member.GetChatMemberPushConfigListResp
	)
	resp = s.chatMemberClient.GetChatMemberPushConfigList(userListReq)
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

func (s *pushService) chatMessagePush(req *pb_gw.OnlinePushMessageReq, serverId int32) {
	var (
		pushResp *pb_gw.OnlinePushMessageResp
	)
	//TODO:serverId分组推送
	pushResp = s.messageGatewayClient.OnlinePushMessage(req)
	if pushResp == nil {
		xlog.Warn(ERROR_CODE_PUSH_GRPC_SERVICE_FAILURE, ERROR_PUSH_GRPC_SERVICE_FAILURE)
		return
	}
	return
}
