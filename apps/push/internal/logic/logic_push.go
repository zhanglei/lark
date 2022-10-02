package logic

import (
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_push"
	"lark/pkg/utils"
)

type OnlinePushMessageHandler func(req *pb_gw.OnlinePushMessageReq, serverId int32)

type PushLogic struct {
	req     *pb_push.PushMessageReq
	handler OnlinePushMessageHandler
}

func PushMessageToHashmapMembers(req *pb_push.PushMessageReq, hashmap map[string]string, handler OnlinePushMessageHandler) {
	pushMessageHandler(req, handler, true, hashmap, nil)
}

func PushMessageToMembers(req *pb_push.PushMessageReq, configs []*pb_chat_member.ChatMemberPushConfig, handler OnlinePushMessageHandler) {
	pushMessageHandler(req, handler, false, nil, configs)
}

func pushMessageHandler(req *pb_push.PushMessageReq, handler OnlinePushMessageHandler, isHash bool, hashmap map[string]string, configs []*pb_chat_member.ChatMemberPushConfig) {
	var (
		length int
	)
	if isHash == true {
		length = len(hashmap)
	} else {
		length = len(configs)
	}
	if length == 0 {
		return
	}
	var (
		logic *PushLogic
	)
	logic = &PushLogic{
		req:     req,
		handler: handler,
	}
	if isHash {
		logic.groupFromHashmap(hashmap)
	} else {
		logic.groupFromConfigs(configs)
	}
}

func (logic *PushLogic) groupFromHashmap(hashmap map[string]string) {
	var (
		jsonStr       string
		member        *pb_gw.OnlinePushMember
		serverMembers = make(map[int32][]*pb_gw.OnlinePushMember)
	)
	for _, jsonStr = range hashmap {
		member = new(pb_gw.OnlinePushMember)
		utils.Unmarshal(jsonStr, member)
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], member)
	}
	logic.pushMessage(serverMembers)
}

func (logic *PushLogic) groupFromConfigs(configs []*pb_chat_member.ChatMemberPushConfig) {
	var (
		conf          *pb_chat_member.ChatMemberPushConfig
		member        *pb_gw.OnlinePushMember
		serverMembers = make(map[int32][]*pb_gw.OnlinePushMember)
	)
	for _, conf = range configs {
		member = &pb_gw.OnlinePushMember{
			Uid:      conf.Uid,
			ServerId: conf.ServerId,
			Platform: conf.Platform,
			Mute:     conf.Mute,
		}
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], member)
	}
	logic.pushMessage(serverMembers)
}

func (logic *PushLogic) pushMessage(serverMembers map[int32][]*pb_gw.OnlinePushMember) {
	var (
		serverId int32
		members  []*pb_gw.OnlinePushMember
		req      *pb_gw.OnlinePushMessageReq
	)
	for serverId, members = range serverMembers {
		if len(members) == 0 {
			continue
		}
		req = &pb_gw.OnlinePushMessageReq{
			Topic:    logic.req.Topic,
			SubTopic: logic.req.SubTopic,
			Msg:      logic.req.Msg,
			Members:  members,
		}
		logic.handler(req, serverId)
	}
}
