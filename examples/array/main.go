package main

import (
	"fmt"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_obj"
	"lark/pkg/utils"
	"strings"
	"time"
)

func main() {
	Test02()
}

func Test03() {
	var (
		count   = 10000
		list    = make([][]int64, count)
		i       int
		jsonStr string
		j       int
		array   [][]int64
	)
	count = 10000

	for i = 0; i < 10000; i++ {
		list[i] = []int64{1, 2, 3, 4}
	}
	jsonStr, _ = utils.Marshal(list)

	for j = 0; j < 100; j++ {
		fmt.Println("开始")
		start := time.Now()
		array = make([][]int64, 0)
		utils.Unmarshal(jsonStr, &array)
		fmt.Println("耗时(毫秒)", time.Now().Sub(start).Milliseconds())
	}
}

func Test02() {
	var (
		str string
		j   int
	)
	str = "1,2,3,4"
	fmt.Println("开始")
	for j = 0; j < 100; j++ {
		var (
			start   = time.Now()
			count   = 10000
			members = make([]*pb_obj.Int64Array, count)
			i       = 0
		)
		for i = 0; i < count; i++ {
			members[i], _ = getArray(str)
		}
		fmt.Println("耗时(毫秒)", time.Now().Sub(start).Milliseconds())
	}
}

func getArray(str string) (array *pb_obj.Int64Array, serverId int64) {
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

func Test01() {
	var (
		m       *pb_chat_member.PushMember
		jsonStr string
		j       int
	)
	m = &pb_chat_member.PushMember{
		Uid:      1,
		ServerId: 1,
		Platform: 2,
		Mute:     1,
	}
	jsonStr, _ = utils.Marshal(m)
	fmt.Println("开始")
	for j = 0; j < 100; j++ {
		var (
			start   = time.Now()
			count   = 10000
			members = make([]*pb_chat_member.PushMember, count)
			i       = 0
		)
		for i = 0; i < count; i++ {
			member := &pb_chat_member.PushMember{}
			utils.Unmarshal(jsonStr, &member)
			members[i] = member
		}
		fmt.Println("耗时(毫秒)", time.Now().Sub(start).Milliseconds())
	}
}
