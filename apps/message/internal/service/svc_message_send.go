package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/protocol"
	"lark/pkg/utils"
)

func setSendChatMessageResp(resp *pb_msg.SendChatMessageResp, code int32, msg string) {
	resp.Code = code
	resp.Msg = msg
}

func (s *messageService) SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error) {
	resp = new(pb_msg.SendChatMessageResp)
	var (
		msg   = new(protocol.ChatMessage)
		inbox = &pb_mq.InboxMessage{
			Topic:    req.Topic,
			SubTopic: req.SubTopic,
			Msg:      new(pb_msg.SrvChatMessage),
		}
		//memberInfo *pb_chat_member.ChatMemberInfo
		ok    bool
		seqId int64
		key   string
		err   error
	)
	// 1、重复消息校验
	key = constant.RK_MSG_CLI_MSG_ID + utils.Int64ToStr(req.Msg.ChatId) + ":" + utils.Int64ToStr(req.Msg.CliMsgId)
	if xredis.KeyExists(key) == true {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_REPEATED_MESSAGE, ERROR_MESSAGE_REPEATED_MESSAGE)
		xlog.Warn(ERROR_CODE_MESSAGE_REPEATED_MESSAGE, ERROR_MESSAGE_REPEATED_MESSAGE)
		return
	}
	if err = xredis.Set(key, req.Msg.CliMsgId, constant.CONST_DURATION_MSG_ID_SECOND); err != nil {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_RECORD_MSGID_FAILED, ERROR_MESSAGE_RECORD_MSGID_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_RECORD_MSGID_FAILED, ERROR_MESSAGE_RECORD_MSGID_FAILED, err.Error())
		return
	}
	// 2、参数校验
	copier.Copy(msg, req.Msg)
	if err = s.validate.Struct(msg); err != nil {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR)
		xlog.Warn(ERROR_CODE_MESSAGE_VALIDATOR_ERR, ERROR_MESSAGE_VALIDATOR_ERR, err.Error())
		return
	}
	// 3、校验是否可以发送消息 & 获取发送人信息
	if ok = s.verifySession(msg.ChatId, pb_enum.CHAT_TYPE(msg.ChatType), msg.SenderId, msg.ReceiverId); ok == false {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_VERIFY_IDENTITY_FAILED, ERROR_MESSAGE_VERIFY_IDENTITY_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_VERIFY_IDENTITY_FAILED, ERROR_MESSAGE_VERIFY_IDENTITY_FAILED)
		return
	}
	// 4、补充消息内容
	if seqId, err = xredis.IncrSeqID(req.Msg.ChatId); err != nil {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_GET_SEQ_ID_FAILED, ERROR_MESSAGE_GET_SEQ_ID_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_GET_SEQ_ID_FAILED, ERROR_MESSAGE_GET_SEQ_ID_FAILED, err.Error())
		return
	}
	copier.Copy(inbox.Msg, req.Msg)
	inbox.Msg.SrvMsgId = xsnowflake.NewSnowflakeID()
	inbox.Msg.SeqId = seqId
	inbox.Msg.SrvTs = utils.NowMilli()
	inbox.Msg.MsgFrom = pb_enum.MSG_FROM_USER
	if req.Msg.ChatType == pb_enum.CHAT_TYPE_GROUP {
		inbox.Msg.ReceiverId = 0
	}

	// 5、将消息推送到kafka消息队列
	_, _, err = s.producer.EnQueue(inbox, utils.Int64ToStr(inbox.Msg.ChatId))
	if err != nil {
		setSendChatMessageResp(resp, ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED)
		xlog.Warn(ERROR_CODE_MESSAGE_ENQUEUE_FAILED, ERROR_MESSAGE_ENQUEUE_FAILED, err.Error())
		return
	}
	return
}

func (s *messageService) verifySession(chatId int64, chatType pb_enum.CHAT_TYPE, senderId int64, receiverId int64) (ok bool) {
	var (
		key     string
		uidList []int64
		list    []interface{}
		req     *pb_chat_member.ChatMemberVerifyReq
		resp    *pb_chat_member.ChatMemberVerifyResp
	)
	key = constant.RK_SYNC_CHAT_MEMBERS_INFO_HASH + utils.Int64ToStr(chatId)
	switch chatType {
	case pb_enum.CHAT_TYPE_PRIVATE:
		list = xredis.HMGet(key, utils.Int64ToStr(receiverId))
		if len(list) == 1 && list[0] != nil {
			ok = true
			return
		}
		uidList = []int64{senderId, receiverId}
	case pb_enum.CHAT_TYPE_GROUP:
		list = xredis.HMGet(key, utils.Int64ToStr(senderId))
		if len(list) == 1 && list[0] != nil {
			ok = true
			return
		}
		uidList = []int64{senderId}
	}

	req = &pb_chat_member.ChatMemberVerifyReq{
		ChatId:   chatId,
		ChatType: chatType,
		UidList:  uidList,
	}

	resp = s.chatMemberClient.ChatMemberVerify(req)
	if resp == nil {
		xlog.Warn(ERROR_CODE_MESSAGE_GRPC_SERVICE_FAILURE, ERROR_MESSAGE_GRPC_SERVICE_FAILURE)
		return
	}
	if resp.MemberInfo != nil && resp.MemberInfo.Uid > 0 {
		ok = true
	}
	return
}
