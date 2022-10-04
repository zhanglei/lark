package repos

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"lark/domain/pos"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
)

type MessageHotRepository interface {
	Create(message *pos.Message) (err error)
}

type messageHotRepository struct {
}

func NewMessageHotRepository() MessageHotRepository {
	return &messageHotRepository{}
}

func (r *messageHotRepository) Create(message *pos.Message) (err error) {
	var (
		coll   *mongo.Collection
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel, coll = entity.Collection(pos.MongoCollectionMessages)
	defer cancel()
	if coll == nil {
		return
	}
	if _, err = coll.InsertOne(ctx, message); err != nil {
		xlog.Error(err.Error())
		return
	}
	return
}

func (r *messageHotRepository) Messages(w *entity.MongoWhere) (messages []*pos.Message, err error) {
	messages = make([]*pos.Message, 0)
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
	err = cur.All(ctx, &messages)
	return
}
