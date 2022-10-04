package service

import (
	"context"
	"github.com/jinzhu/copier"
	"lark/domain/po"
	"lark/pkg/common/xredis"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
	"sort"
)

func (s *chatMessageService) GetChatMessages(_ context.Context, req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp, _ error) {
	resp = &pb_chat_msg.GetChatMessagesResp{List: make([]*pb_msg.SrvChatMessage, 0)}
	var (
		nowTs       = utils.NowMilli()
		list        = make([]*po.Message, 0)
		cacheList   []*po.Message
		hotList     []*po.Message
		historyList []*po.Message
		msgCount    int
		maxSeqId    uint64
		next        bool
		err         error
	)
	// 1、消息边界
	maxSeqId, _ = xredis.GetMaxSeqID(req.ChatId)
	if req.SeqId >= int64(maxSeqId) {
		if req.New == true {
			// 1.1 消息越界
			return
		}
		req.SeqId = int64(maxSeqId)
	}

	if nowTs-req.MsgTs < s.conf.MsgCache.L1Duration {
		// 2、从redis缓存中读取
		cacheList, next, err = s.GetCacheMessages(req, int64(maxSeqId))
		if next == false || err != nil {
			if len(cacheList) > 0 {
				if req.New == false {
					sortMessageList(cacheList, true)
				}
				copier.Copy(&resp.List, cacheList)
			}
			return
		}
	}

	if nowTs-req.MsgTs < s.conf.MsgCache.L2Duration {
		// 3、从mongo缓存中读取
		hotList, next, err = s.GetHotMessages(req, int64(maxSeqId))
		if next == false || err != nil {
			if len(cacheList) > 0 {
				hotList = append(cacheList, hotList...)
				if req.New == false {
					sortMessageList(hotList, true)
				}
			}
			if len(hotList) > 0 {
				copier.Copy(&resp.List, hotList)
			}
			return
		}
	}

	// 4、从mysql缓存中读取
	historyList, err = s.GetHistoryMessages(req)
	if err != nil {
		return
	}

	if len(cacheList) > 0 {
		list = append(list, cacheList...)
	}
	if len(hotList) > 0 {
		list = append(list, hotList...)
	}
	if len(historyList) > 0 {
		list = append(list, historyList...)
	}

	msgCount = len(list)
	if msgCount == 0 {
		return
	}
	if req.New == false {
		sortMessageList(list, true)
	}
	copier.Copy(&resp.List, list)
	return
}

func sortMessageList(list []*po.Message, asc bool) {
	sort.Slice(list, func(i, j int) bool {
		if asc == true {
			return list[i].SeqId < list[j].SeqId
		} else {
			return list[i].SeqId > list[j].SeqId
		}
	})
}
