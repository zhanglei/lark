package gateway

import (
	"context"
	"google.golang.org/protobuf/proto"
	"lark/apps/msg_gateway/internal/logic"
	"lark/pkg/common/xgopool"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/utils"
)

func setOnlinePushMessageResp(resp *pb_gw.OnlinePushMessageResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *gatewayServer) OnlinePushMessage(ctx context.Context, req *pb_gw.OnlinePushMessageReq) (resp *pb_gw.OnlinePushMessageResp, _ error) {
	resp = &pb_gw.OnlinePushMessageResp{PushResps: make([]*pb_gw.PlatformPushResp, 0)}
	var (
		buf []byte
		err error
	)
	buf, err = proto.Marshal(req.Msg)
	if err != nil {
		setOnlinePushMessageResp(resp, ERROR_CODE_GATEWAY_MSG_MARSHAL_ERR, ERROR_GATEWAY_MSG_MARSHAL_ERR)
		xlog.Warn(ERROR_CODE_GATEWAY_MSG_MARSHAL_ERR, ERROR_GATEWAY_MSG_MARSHAL_ERR, err)
		return
	}
	buf, err = utils.Encode(int32(req.Topic), int32(req.SubTopic), int32(pb_enum.MESSAGE_TYPE_NEW), buf)
	if err != nil {
		setOnlinePushMessageResp(resp, ERROR_CODE_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, ERROR_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR)
		xlog.Warn(ERROR_CODE_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, ERROR_GATEWAY_MSG_ASSEMBLY_PROTOCOL_ERR, err)
		return
	}
	xgopool.Go(func() {
		logic.SendMessage(req, buf, s.wsServer.SendMessage, s.OfflinePushMessage)
	})
	return
}
