package repos

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"lark/domain/pos"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
)

func (r *chatMessageRepository) HotMessages(w *entity.MongoWhere) (list []*pos.Message, err error) {
	list = make([]*pos.Message, 0)
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
		cur    *mongo.Cursor
	)
	ctx, cancel, coll = entity.Collection(pos.MongoCollectionMessages)
	defer cancel()
	if coll == nil {
		return
	}
	cur, err = coll.Find(ctx, w.Filter, w.FindOptions)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &list)
	return
}
