package config

import (
	"github.com/cpyun/gyopls-core/storage"
	"github.com/cpyun/gyopls-core/storage/queue"
	"github.com/robinjoseph08/redisqueue/v2"
)

type Queue struct {
	Redis  *QueueRedis
	Memory *QueueMemory
	NSQ    *QueueNSQ `json:"nsq" yaml:"nsq"`
}

type QueueRedis struct {
	Redis
	Producer *redisqueue.ProducerOptions
	Consumer *redisqueue.ConsumerOptions
}

type QueueMemory struct {
	PoolSize uint
}

type QueueNSQ struct {
	//NSQOptions
	ChannelPrefix string
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil && e.Redis == nil && e.NSQ == nil
}

// Setup 启用顺序 redis > 其他 > memory
func (e Queue) Setup() (storage.AdapterQueue, error) {
	//if e.Redis != nil {
	//	e.Redis.Consumer.ReclaimInterval = e.Redis.Consumer.ReclaimInterval * time.Second
	//	e.Redis.Consumer.BlockingTimeout = e.Redis.Consumer.BlockingTimeout * time.Second
	//	e.Redis.Consumer.VisibilityTimeout = e.Redis.Consumer.VisibilityTimeout * time.Second
	//	client := GetRedisClient()
	//	if client == nil {
	//		options, err := e.Redis.Redis.GetRedisOptions()
	//		if err != nil {
	//			return nil, err
	//		}
	//		client = redis.NewClient(options)
	//		_redis = client
	//	}
	//	e.Redis.Producer.RedisClient = client
	//	e.Redis.Consumer.RedisClient = client
	//	return queue.NewRedis(e.Redis.Producer, e.Redis.Consumer)
	//}
	//if e.NSQ != nil {
	//	cfg, err := e.NSQ.GetNSQOptions()
	//	if err != nil {
	//		return nil, err
	//	}
	//	return queue.NewNSQ(e.NSQ.Addresses, cfg, e.NSQ.ChannelPrefix)
	//}
	return queue.NewMemory(e.Memory.PoolSize), nil
}
