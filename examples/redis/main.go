package main

import (
	"lark/examples/config"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/utils"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	var maps = map[string]interface{}{}
	for i := 0; i < 10; i++ {
		p := pb_gw.OnlinePushMember{
			Uid:      int64(i + 1),
			ServerId: 8,
			Platform: 2,
			Mute:     1,
		}
		maps[utils.Int64ToStr(p.Uid)], _ = utils.Marshal(p)
	}
	err := xredis.HMSet("MEMBERS:2", maps)
	if err != nil {
		xlog.Warn(err.Error())
	}
}
