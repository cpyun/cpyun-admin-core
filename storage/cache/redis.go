package cache

import (
	"context"
	"errors"
	"github.com/cpyun/cpyun-admin-core/storage"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis cache implement
type Redis struct {
	ctx    context.Context
	client *redis.Client
	prefix string
}

// NewRedis redis模式
func NewRedis(client *redis.Client, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = redis.NewClient(options)
	}

	ctx := context.TODO()
	r := &Redis{
		ctx:    ctx,
		client: client,
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(r.ctx).Result()
	return err
}

func (r *Redis) SetPrefix(s string) {
	r.prefix = s
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	key = r.prefix + key
	return r.client.Get(r.ctx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	key = r.prefix + key
	return r.client.Set(r.ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	key = r.prefix + key
	return r.client.Del(r.ctx, key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	key = r.prefix + key
	return r.client.HGet(r.ctx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	key = r.prefix + key
	return r.client.HDel(r.ctx, hk, key).Err()
}

// Increase value
func (r *Redis) Increase(key string) error {
	key = r.prefix + key
	return r.client.Incr(r.ctx, key).Err()
}

func (r *Redis) Decrease(key string) error {
	key = r.prefix + key
	return r.client.Decr(r.ctx, key).Err()
}

// Expire Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	key = r.prefix + key
	return r.client.Expire(r.ctx, key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func CovertInterfaceToStruct(face storage.AdapterCache) (*Redis, error) {
	value := reflect.ValueOf(face)
	if value.IsNil() {
		return nil, errors.New("value is nil")
	} else if value.Kind() != reflect.Ptr {
		return nil, errors.New("error of kind [pointer]")
	}

	//// 取数据
	//value = value.Elem()
	//if value.Kind() != reflect.Struct {
	//	return minio, errors.New("not a struct")
	//}

	return value.Interface().(*Redis), nil
}
