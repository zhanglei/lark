package logic

import (
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/utils"
)

func GetOnlinePushMembersFromHash(hashmap map[string]string) (serverMembers map[int32][]*pb_chat_member.PushMember) {
	return pushMessageHandler(true, hashmap, nil)
}

func GetOnlinePushMembersFromList(members []*pb_chat_member.PushMember) (serverMembers map[int32][]*pb_chat_member.PushMember) {
	return pushMessageHandler(false, nil, members)
}

func pushMessageHandler(isHash bool, hashmap map[string]string, members []*pb_chat_member.PushMember) (serverMembers map[int32][]*pb_chat_member.PushMember) {
	var (
		length int
	)
	if isHash == true {
		length = len(hashmap)
	} else {
		length = len(members)
	}
	if length == 0 {
		return
	}

	if isHash {
		return groupFromHashmap(hashmap)
	} else {
		return groupFromMembers(members)
	}
}

func groupFromHashmap(hashmap map[string]string) (serverMembers map[int32][]*pb_chat_member.PushMember) {
	var (
		jsonStr string
		member  *pb_chat_member.PushMember
	)
	serverMembers = make(map[int32][]*pb_chat_member.PushMember)
	for _, jsonStr = range hashmap {
		member = new(pb_chat_member.PushMember)
		utils.Unmarshal(jsonStr, member)
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], member)
	}
	return
}

func groupFromMembers(members []*pb_chat_member.PushMember) (serverMembers map[int32][]*pb_chat_member.PushMember) {
	var (
		member *pb_chat_member.PushMember
		index  int
	)
	serverMembers = make(map[int32][]*pb_chat_member.PushMember)
	for index, member = range members {
		serverMembers[member.ServerId] = append(serverMembers[member.ServerId], members[index])
	}
	return
}
