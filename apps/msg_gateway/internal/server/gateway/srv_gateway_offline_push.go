package gateway

import (
	"lark/pkg/proto/pb_ofps"
	"lark/pkg/utils"
)

func (s *gatewayServer) OfflinePushMessage(req *pb_ofps.OfflinePushMessageReq) {
	s.producer.EnQueue(req, utils.Int64ToStr(req.ChatId))
}
