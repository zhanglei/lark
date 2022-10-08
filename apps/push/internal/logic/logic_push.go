package logic

import (
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/utils"
)

func GetOnlinePushMembersFromHash(hashmap map[string]string) (serverMembers map[int32][]*pb_gw.OnlinePushMember) {
	return pushMessageHandler(true, hashmap, nil)
}

func GetOnlinePushMembersFromList(configs []*pb_chat_member.ChatMemberPushConfig) (serverMembers map[int32][]*pb_gw.OnlinePushMember) {
	return pushMessageHandler(false, nil, configs)
}

func pushMessageHandler(isHash bool, hashmap map[string]string, configs []*pb_chat_member.ChatMemberPushConfig) (serverMembers map[int32][]*pb_gw.OnlinePushMember) {
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

	if isHash {
		return groupFromHashmap(hashmap)
	} else {
		return groupFromConfigs(configs)
	}
}

func groupFromHashmap(hashmap map[string]string) (serverMembers map[int32][]*pb_gw.OnlinePushMember) {
	var (
		jsonStr string
		member  *pb_gw.OnlinePushMember
	)
	serverMembers = make(map[int32][]*pb_gw.OnlinePushMember)
	for _, jsonStr = range hashmap {
		member = new(pb_gw.OnlinePushMember)
		utils.Unmarshal(jsonStr, member)
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], member)
	}
	return
}

func groupFromConfigs(configs []*pb_chat_member.ChatMemberPushConfig) (serverMembers map[int32][]*pb_gw.OnlinePushMember) {
	var (
		conf   *pb_chat_member.ChatMemberPushConfig
		member *pb_gw.OnlinePushMember
	)
	serverMembers = make(map[int32][]*pb_gw.OnlinePushMember)
	for _, conf = range configs {
		member = &pb_gw.OnlinePushMember{
			Uid:      conf.Uid,
			ServerId: conf.ServerId,
			Platform: conf.Platform,
			Mute:     conf.Mute,
		}
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], member)
	}
	return
}
