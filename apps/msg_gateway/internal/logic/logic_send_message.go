package logic

import (
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_obj"
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

func (logic *SendMessageLogic) push(members []*pb_obj.Int64Array, msgBuf []byte) {
	var (
		member      *pb_obj.Int64Array
		result      int32
		ofpsMembers = make([]*pb_ofps.OfflinePushMember, 0)
		uid         int64
		platform    pb_enum.PLATFORM_TYPE
		mute        pb_enum.MUTE_TYPE
	)
	//member.Uid, member.Platform, member.ServerId, member.Mute
	for _, member = range members {
		uid = member.Vals[0]
		platform = pb_enum.PLATFORM_TYPE(member.Vals[1])
		mute = pb_enum.MUTE_TYPE(member.Vals[3])

		result = logic.onlinePushHandler(uid, int32(platform), msgBuf)
		if mute == pb_enum.MUTE_TYPE_OPENED {
			continue
		}
		if uid == logic.req.Msg.SenderId {
			continue
		}
		if result > 0 {
			ofpsMember := &pb_ofps.OfflinePushMember{
				Uid:      uid,
				Platform: platform,
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
