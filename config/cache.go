package config

import (
	"github.com/cpyun/cpyun-admin-core/sdk"
	"github.com/cpyun/cpyun-admin-core/storage"
	"github.com/cpyun/cpyun-admin-core/storage/cache"
)

type Cache struct {
	Driver string `yaml:"driver"`
	Redis  *Redis
	Memory interface{}
}

var CacheConfig = new(Cache)

// Setup 构造cache 顺序 redis > 其他 > memory
func (e *Cache) Setup() (storage.AdapterCache, error) {
	var cacheAdapter storage.AdapterCache

	if e.Driver == "redis" {
		//e.Redis = &Settings.Redis

		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}

		client := GetRedisClient()
		cacheAdapter, err = cache.NewRedis(client, options)
		if err != nil {
			return nil, err
		}
	} else {
		cacheAdapter = cache.NewMemory()
	}

	sdk.Runtime.SetCacheAdapter(cacheAdapter)

	return cacheAdapter, nil
}
