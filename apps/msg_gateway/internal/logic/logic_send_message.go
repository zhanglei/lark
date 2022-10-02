package logic

import (
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_ofps"
)

type OnlinePushMessageHandler func(uid int64, platform int32, message []byte) (result int32)
type OfflinePushMessageHandler func(req *pb_ofps.OfflinePushMessageReq)

type SendMessageLogic struct {
	req                *pb_gw.OnlinePushMessageReq
	onlinePushHandler  OnlinePushMessageHandler
	offlinePushHandler OfflinePushMessageHandler
}

func SendMessage(req *pb_gw.OnlinePushMessageReq, msgBuf []byte, onlinePushHandler OnlinePushMessageHandler, offlinePushHandler OfflinePushMessageHandler) {
	if len(req.Members) == 0 {
		return
	}
	if len(msgBuf) == 0 {
		return
	}
	var (
		logic *SendMessageLogic
	)
	logic = &SendMessageLogic{
		req:                req,
		onlinePushHandler:  onlinePushHandler,
		offlinePushHandler: offlinePushHandler,
	}
	logic.push(req.Members, msgBuf)
}

func (logic *SendMessageLogic) push(members []*pb_gw.OnlinePushMember, msgBuf []byte) {
	var (
		member      *pb_gw.OnlinePushMember
		result      int32
		ofpsMembers = make([]*pb_ofps.OfflinePushMember, 0)
	)
	for _, member = range members {
		result = logic.onlinePushHandler(member.Uid, int32(member.Platform), msgBuf)
		if member.Mute == pb_enum.MUTE_TYPE_OPENED {
			continue
		}
		if result > 0 {
			ofpsMember := &pb_ofps.OfflinePushMember{
				Uid:      member.Uid,
				Platform: member.Platform,
			}
			ofpsMembers = append(ofpsMembers, ofpsMember)
		}
	}
	logic.offlinePush(ofpsMembers)
}

func (logic *SendMessageLogic) offlinePush(members []*pb_ofps.OfflinePushMember) {
	req := &pb_ofps.OfflinePushMessageReq{
		Topic:    logic.req.Topic,
		SubTopic: logic.req.SubTopic,
		ChatId:   logic.req.Msg.ChatId,
		Member:   members,
		MsgType:  0,
		Body:     nil,
	}
	logic.offlinePushHandler(req)
}
