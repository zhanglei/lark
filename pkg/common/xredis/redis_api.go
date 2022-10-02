package xredis

import (
	"context"
	"strconv"
	"time"
)

func RealKey(key string) string {
	if cli != nil {
		return cli.Prefix + key
	}
	return key
}

func Del(key string) error {
	key = RealKey(key)
	return cli.client.Del(context.Background(), key).Err()
}

func KeyExists(key string) (ok bool) {
	key = RealKey(key)
	val := cli.client.Exists(context.Background(), key).Val()
	if val == 1 {
		ok = true
	}
	return
}

func Set(key string, value interface{}, expire time.Duration) error {
	key = RealKey(key)
	if expire > 0 {
		return cli.client.Set(context.Background(), key, value, expire).Err()
	}
	return cli.client.Set(context.Background(), key, value, 0).Err()
}

func Get(key string) (string, error) {
	key = RealKey(key)
	return cli.client.Get(context.Background(), key).Result()
}

func Incr(key string) (int64, error) {
	key = RealKey(key)
	return cli.client.Incr(context.Background(), key).Result()
}

func GetUint64(key string) (uint64, error) {
	key = RealKey(key)
	return cli.client.Get(context.Background(), key).Uint64()
}

func GetInt(key string) (int, error) {
	key = RealKey(key)
	return cli.client.Get(context.Background(), key).Int()
}

func HGetInt(key, field string) (value int, err error) {
	key = RealKey(key)
	return cli.client.HGet(context.Background(), key, field).Int()
}

func HGetAll(key string) map[string]string {
	key = RealKey(key)
	hash := cli.client.HGetAll(context.Background(), key).Val()
	return hash
}

func HSet(key string, value interface{}) error {
	key = RealKey(key)
	return cli.client.HSet(context.Background(), key, value).Err()
}

func HSetNX(key, field string, value interface{}) error {
	key = RealKey(key)
	return cli.client.HSetNX(context.Background(), key, field, value).Err()
}

func HDels(key string, fields []string) error {
	key = RealKey(key)
	return cli.client.HDel(context.Background(), key, fields...).Err()
}

func HDel(key string, field string) error {
	key = RealKey(key)
	return cli.client.HDel(context.Background(), key, field).Err()
}

func HMSet(key string, values map[string]interface{}) error {
	key = RealKey(key)
	return cli.client.HMSet(context.Background(), key, values).Err()
}

func HMGet(key string, fields ...string) []interface{} {
	key = RealKey(key)
	return cli.client.HMGet(context.Background(), key, fields...).Val()
}

// Sequence ID
func GetMaxSeqID(chatId int64) (uint64, error) {
	key := MSG_SEQ_ID + strconv.FormatInt(chatId, 10)
	return GetUint64(key)
}

func IncrSeqID(chatId int64) (int64, error) {
	key := MSG_SEQ_ID + strconv.FormatInt(chatId, 10)
	return Incr(key)
}

func SAdd(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return cli.client.SAdd(context.Background(), key, members).Err()
}

func Srem(key string, members ...interface{}) (err error) {
	key = RealKey(key)
	return cli.client.SRem(context.Background(), key, members).Err()
}

func Smembers(key string) []string {
	key = RealKey(key)
	return cli.client.SMembers(context.Background(), key).Val()
}
