package service

import (
	"lark/domain/pos"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/utils"
	"strconv"
)

func (s *chatMessageService) GetCacheMessages(req *pb_chat_msg.GetChatMessagesReq, maxSeqId int64) (list []*pos.Message, next bool, err error) {
	var (
		msgCount int
	)
	list = s.GetCacheChatMessages(req, int(maxSeqId))
	msgCount = len(list)
	if msgCount == int(req.Limit) {
		return
	}
	if msgCount > 0 {
		if req.New == true {
			req.SeqId = list[msgCount-1].SeqId
		} else {
			req.SeqId = list[0].SeqId
		}
		if req.SeqId >= maxSeqId {
			return
		}
		req.Limit -= int32(msgCount)
	}
	next = true
	return
}

func (s *chatMessageService) GetCacheChatMessages(req *pb_chat_msg.GetChatMessagesReq, max int) (list []*pos.Message) {
	list = make([]*pos.Message, 0)
	var (
		minSeqId int
		maxSeqId int
		seqId    int
		dv       int
		index    int
		key      string
		jsonStr  string
		err      error
	)
	if req.New == true {
		minSeqId = int(req.SeqId) + 1
		maxSeqId = minSeqId + int(req.Limit)
	} else {
		maxSeqId = int(req.SeqId) - 1
		minSeqId = maxSeqId - int(req.Limit)
	}
	if minSeqId < 0 {
		minSeqId = 0
	}
	if maxSeqId > max {
		maxSeqId = max
	}
	dv = maxSeqId - minSeqId
	for index = 0; index < dv; index++ {
		if req.New == true {
			seqId = minSeqId + index
		} else {
			seqId = maxSeqId - index
		}
		key = constant.RK_MSG_CACHE + utils.Int64ToStr(req.ChatId) + ":" + strconv.Itoa(seqId)
		jsonStr, err = xredis.Get(key)
		if err != nil || jsonStr == "" {
			break
		}
		var msg pos.Message
		utils.Unmarshal(jsonStr, &msg)
		list = append(list, &msg)
	}
	return
}

func (s *chatMessageService) SaveCacheChatMessageCache(list []*pos.Message) {
	if len(list) == 0 {
		return
	}
	var (
		msg     *pos.Message
		jsonStr string
		key     string
	)
	for _, msg = range list {
		jsonStr, _ = utils.Marshal(msg)
		if jsonStr == "" {
			continue
		}
		key = constant.RK_MSG_CACHE + msg.GetChatId() + ":" + msg.GetSeqId()
		xredis.Set(key, jsonStr, constant.CONST_DURATION_MSG_CACHE_SECOND)
	}
}
