package logic

import (
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"strings"
)

func GetOnlinePushMembersFromHash(hashmap map[string]string) (serverMembers map[int64][]*pb_obj.Int64Array) {
	return pushMessageHandler(true, hashmap, nil)
}

func GetOnlinePushMembersFromList(members []string) (serverMembers map[int64][]*pb_obj.Int64Array) {
	return pushMessageHandler(false, nil, members)
}

func pushMessageHandler(isHash bool, hashmap map[string]string, members []string) (serverMembers map[int64][]*pb_obj.Int64Array) {
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

func groupFromHashmap(hashmap map[string]string) (serverMembers map[int64][]*pb_obj.Int64Array) {
	var (
		str      string
		array    *pb_obj.Int64Array
		serverId int64
	)
	serverMembers = make(map[int64][]*pb_obj.Int64Array)
	for _, str = range hashmap {
		array, serverId = getMemberInfo(str)
		if array == nil {
			continue
		}
		serverMembers[serverId] = append(serverMembers[serverId], array)
	}
	return
}

func groupFromMembers(members []string) (serverMembers map[int64][]*pb_obj.Int64Array) {
	var (
		str      string
		array    *pb_obj.Int64Array
		serverId int64
	)
	serverMembers = make(map[int64][]*pb_obj.Int64Array)
	for _, str = range members {
		array, serverId = getMemberInfo(str)
		if array == nil {
			continue
		}
		serverMembers[serverId] = append(serverMembers[serverId], array)
	}
	return
}

func getMemberInfo(str string) (array *pb_obj.Int64Array, serverId int64) {
	var (
		arr []string
	)
	arr = strings.Split(str, ",")
	if len(arr) != 4 {
		return
	}
	array = &pb_obj.Int64Array{Vals: make([]int64, 4)}

	// member.Uid, member.Platform, member.ServerId, member.Mute
	array.Vals[0], _ = utils.ToInt64(arr[0])
	array.Vals[1], _ = utils.ToInt64(arr[1])
	serverId, _ = utils.ToInt64(arr[2])
	array.Vals[2] = serverId
	array.Vals[3], _ = utils.ToInt64(arr[3])
	return
}
