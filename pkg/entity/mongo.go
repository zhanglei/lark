package entity

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lark/pkg/common/xmongo"
	"time"
)

func Collection(collection string) (ctx context.Context, cancel context.CancelFunc, coll *mongo.Collection) {
	var (
		db *mongo.Database
	)
	ctx, cancel = NewContext()
	db = xmongo.GetDB()
	if db == nil {
		return
	}
	coll = db.Collection(collection)
	return
}

func NewContext() (ctx context.Context, cancelFunc context.CancelFunc) {
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	return
}

type MongoWhere struct {
	Filter      map[string]interface{}
	FindOptions *options.FindOptions
}

func NewMongoWhere() *MongoWhere {
	return &MongoWhere{
		Filter:      make(map[string]interface{}),
		FindOptions: new(options.FindOptions),
	}
}

func (m *MongoWhere) SetSort(key string, asc bool) {
	var val = 1
	if asc == false {
		val = -1
	}
	m.FindOptions.SetSort(bson.D{bson.E{key, val}})
}

func (m *MongoWhere) SetLimit(limit int64) {
	m.FindOptions.SetLimit(limit)
}

func (m *MongoWhere) SetFilter(key string, val interface{}) {
	m.Filter[key] = val
}
