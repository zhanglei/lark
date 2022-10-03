package service

import (
	"google.golang.org/protobuf/proto"
	"lark/apps/msg_gateway/internal/server/websocket/ws"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
)

func (s *wsService) replyMessage(client *ws.Client, topic pb_enum.TOPIC, subTopic pb_enum.SUB_TOPIC, msgId int64, code int32, msg string) {
	var (
		resp *pb_msg.MessageResp
		buf  []byte
		err  error
	)
	resp = &pb_msg.MessageResp{
		Code:  code,
		Msg:   msg,
		MsgId: msgId,
	}
	buf, err = proto.Marshal(resp)
	if err != nil {
		xlog.Warn(ERROR_CODE_WS_PROTOCOL_MARSHAL_ERR, ERROR_WS_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	buf, err = utils.Encode(int32(topic), int32(subTopic), int32(pb_enum.MESSAGE_TYPE_RESP), buf)
	if err != nil {
		xlog.Warn(ERROR_CODE_WS_ASSEMBLY_PROTOCOL_ERR, ERROR_WS_ASSEMBLY_PROTOCOL_ERR, err.Error())
		return
	}
	client.Send(buf)
}

func (s *wsService) MessageCallback(msg *ws.Message) {
	var (
		req *pb_msg.Packet
	)
	req, _ = utils.Decode(msg.Body)
	switch req.Topic {
	case pb_enum.TOPIC_CHAT:
		s.dispatchMessage(msg.Client, req)
	default:
		s.replyMessage(msg.Client,
			req.Topic,
			req.SubTopic,
			0,
			ERROR_CODE_WS_TOPIC_ID_ERR,
			ERROR_WS_TOPIC_ID_ERR)
		xlog.Warn(ERROR_CODE_WS_TOPIC_ID_ERR, ERROR_WS_TOPIC_ID_ERR, req.Topic)
	}
}

func (s *wsService) dispatchMessage(client *ws.Client, req *pb_msg.Packet) {
	var (
		chatMessageReq *pb_msg.SendChatMessageReq
		resp           *pb_msg.SendChatMessageResp
		err            error
	)
	chatMessageReq = &pb_msg.SendChatMessageReq{
		Topic:    req.Topic,
		SubTopic: req.SubTopic,
		Msg:      new(pb_msg.CliChatMessage),
	}
	err = proto.Unmarshal(req.Data, chatMessageReq.Msg)
	if err != nil {
		s.replyMessage(client,
			req.Topic,
			req.SubTopic,
			0,
			ERROR_CODE_WS_GRPC_SERVICE_FAILURE,
			ERROR_WS_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_WS_GRPC_SERVICE_FAILURE, ERROR_WS_GRPC_SERVICE_FAILURE, err.Error())
		return
	}
	resp = s.msgClient.SendChatMessage(chatMessageReq)
	if resp == nil {
		s.replyMessage(client,
			chatMessageReq.Topic,
			chatMessageReq.SubTopic,
			chatMessageReq.Msg.CliMsgId,
			ERROR_CODE_WS_GRPC_SERVICE_FAILURE,
			ERROR_WS_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_WS_GRPC_SERVICE_FAILURE, ERROR_WS_GRPC_SERVICE_FAILURE)
		return
	}
	s.replyMessage(client,
		chatMessageReq.Topic,
		chatMessageReq.SubTopic,
		chatMessageReq.Msg.CliMsgId,
		resp.Code,
		resp.Msg)
	if resp.Code > 0 {
		xlog.Warn(resp.Code, resp.Msg)
		return
	}
}
