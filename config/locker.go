package config

import (
	"github.com/cpyun/gyopls-core/sdk"
	"github.com/cpyun/gyopls-core/storage"
	"github.com/cpyun/gyopls-core/storage/locker"
	"github.com/redis/go-redis/v9"
)

var LockerConfig = new(Locker)

type Locker struct {
	Driver string `json:"driver" yaml:"driver"`
	Redis  *Redis
	//Memory any
}

func (l Locker) Empty() bool {
	return l.Redis == nil
}

func (l Locker) Setup() (storage.AdapterLocker, error) {
	var lockerApter storage.AdapterLocker

	if l.Driver == "" {
		l.Driver = "redis"
	}

	switch l.Driver {
	case "redis":
		client := GetRedisClient()
		if client == nil {
			options, err := l.Redis.GetRedisOptions()
			if err != nil {
				return nil, err
			}
			client = redis.NewClient(options)
			SetRedisClient(client)
		}

		lockerApter = locker.NewRedis(client)
	}

	sdk.Runtime.SetLockerAdapter(lockerApter)

	return lockerApter, nil
}
