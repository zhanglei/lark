package service

import (
	"lark/domain/pos"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_chat_msg"
)

func (s *chatMessageService) GetHistoryMessages(req *pb_chat_msg.GetChatMessagesReq) (list []*pos.Message, err error) {
	// 从mysql中获取消息
	var (
		w = entity.NewMysqlWhere()
	)
	w.Limit = int(req.Limit)
	w.Query += " AND deleted_ts = 0"
	w.Query += " AND chat_id = ?"
	w.Args = append(w.Args, req.ChatId)
	if req.New {
		w.Sort = "seq_id ASC"
		w.Query += " AND seq_id > ?"
		w.Args = append(w.Args, req.SeqId)
	} else {
		w.Sort = "seq_id DESC"
		w.Query += " AND seq_id < ?"
		w.Args = append(w.Args, req.SeqId)
	}
	list, err = s.chatMessageRepo.HistoryMessages(w)
	if err != nil {
		return
	}
	return
}
