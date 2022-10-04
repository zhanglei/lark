package repo

import (
	"lark/domain/po"
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

func (r *chatMessageRepository) HistoryMessages(w *entity.MysqlWhere) (list []*po.Message, err error) {
	list = make([]*po.Message, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Order(w.Sort).Find(&list).Error
	return
}
