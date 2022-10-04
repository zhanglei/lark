package repos

import (
	"lark/domain/pos"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

func (r *chatMessageRepository) HistoryMessages(w *entity.MysqlWhere) (list []*pos.Message, err error) {
	list = make([]*pos.Message, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Order(w.Sort).Find(&list).Error
	return
}
