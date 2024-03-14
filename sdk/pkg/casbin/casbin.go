package casbin

import (
	"github.com/casbin/casbin/v2/persist"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	redisWatcher "github.com/casbin/redis-watcher/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/cpyun/gyopls-core/config"
	"github.com/cpyun/gyopls-core/logger"
	"github.com/cpyun/gyopls-core/sdk"
)

var (
	enforcer *casbin.SyncedEnforcer
	once     sync.Once
)

func Setup(host string, db *gorm.DB) *casbin.SyncedEnforcer {
	once.Do(func() {
		var err error

		// 获取gorm适配器
		apter, err := gormAdapter.NewAdapterByDB(db)
		if err != nil {
			logger.Fatal("casbin new adapter error: ", err)
		}

		// 根据路径获取Model
		m, err := model.NewModelFromFile(config.CasbinConfig.ModelPath)
		if err != nil {
			logger.Fatal("casbin new model error: ", err)
		}

		enforcer, err = casbin.NewSyncedEnforcer(m, apter)
		if err != nil {
			logger.Fatal("casbin new synced enforcer error: ", err)
		}

		// 设置监视器
		initWatcher()

		err = enforcer.LoadPolicy()
		if err != nil {
			logger.Error("casbin load policy error: ", err)
		}

		log.SetLogger(&Logger{})
		enforcer.EnableLog(true)
	})

	sdk.Runtime.SetCasbin(host, enforcer)
	return enforcer
}

//func initAdapter() {
//	var err error
//
//	switch config.CasbinConfig.Watcher {
//	case "gorm":
//		dbs := sdk.Runtime.GetDb()
//
//		apter, err := gormAdapter.NewAdapterByDB(dbs["*"])
//	case "redis":
//		apter, err := redisAdapter.NewAdapter(casbinCfg.Redis.Addr, redisWatcher.WatcherOptions{})
//	}
//
//	if err != nil {
//		logger.Fatal("casbin new adapter error: ", err)
//	}
//}

// initWatcher 设置监视器
func initWatcher() {
	var err error
	var w persist.Watcher

	casbinCfg := config.CasbinConfig

	switch casbinCfg.Watcher {
	case "redis":
		w, err = redisWatcher.NewWatcher(casbinCfg.Redis.Addr, redisWatcher.WatcherOptions{
			Options: redis.Options{
				Network:  casbinCfg.Redis.Network, // "tcp",
				Password: casbinCfg.Redis.Password,
			},
			Channel:    "/casbin",
			IgnoreSelf: false,
		})
		if err != nil {
			logger.Fatal("casbin new watcher error: ", err)
		}
	default:
		return
	}

	// Set callback
	err = w.SetUpdateCallback(updateCallback)
	if err != nil {
		logger.Fatal("casbin set update callback error: ", err)
	}

	// Set the watcher for the enforcer.
	err = enforcer.SetWatcher(w)
	if err != nil {
		logger.Fatal("casbin set watcher error: ", err)
	}
}

// updateCallback 设置更新回调
func updateCallback(msg string) {
	err := enforcer.LoadPolicy()
	if err != nil {
		logger.Error("casbin update load policy error: ", err)
	}
}
