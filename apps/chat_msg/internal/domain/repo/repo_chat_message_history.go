package repo

import (
	"lark/pkg/common/xmysql"
	"lark/pkg/entity"
)

func (r *chatMessageRepository) HistoryMessages(w *entity.MysqlWhere) (list []*entity.Message, err error) {
	list = make([]*entity.Message, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Limit(w.Limit).Order(w.Sort).Find(&list).Error
	return
}
