package service

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_ofps"
	"sync/atomic"
)

func (s *offlinePushService) MessageHandler(msg []byte, chatId string) (err error) {
	var (
		req = new(pb_ofps.OfflinePushMessageReq)
	)
	proto.Unmarshal(msg, req)
	atomic.AddInt64(&s.msgCount, 1)
	fmt.Println("离线推送:", s.msgCount, len(req.Member))
	return
}
