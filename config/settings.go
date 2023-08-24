package config

import (
	"github.com/cpyun/cpyun-admin-core/config/driver"
	"log"
)

var (
	ExtendConfig interface{}
	Settings     *Config
	_cfg         *settings
)

type settings struct {
	prefix    string // 前缀
	settings  Config
	callbacks []func()
}

// 初始化
func (e *settings) init() {
	// 配置日志
	e.settings.Logger.Setup()
	//配置多数据库
	e.settings.multiDatabase()

	// 调用回调函数
	e.runCallback()
}

// 回调函数
func (e *settings) runCallback() {
	for _, callback := range e.callbacks {
		callback()
	}
}

// OnChange 修改配置
func (e *settings) OnChange() {
	e.init()
	log.Println("!!! config change and reload")
}

// Setup
// @description   Setup 载入配置文件
// @auth      cpYun             时间（2022/7/22   10:00 ）
// @param     s         string        "配置文件路径"
// @param     fs        func          "回调函数"
// @return
func Setup(s string, fs ...func()) {
	Settings = &Config{
		Application: ApplicationConfig,
		Database:    DatabaseConfig,
		Databases:   &DatabasesConfig,
		Filesystem:  FilesystemConfig,
		Cache:       CacheConfig,
		Logger:      LoggerConfig,
		Queue:       QueueConfig,
		Locker:      LockerConfig,
		JWT:         JwtConfig,
		Casbin:      CasbinConfig,
		Extend:      ExtendConfig,
	}

	//1. 读取配置
	driver.WithBindEnv("")
	driver.NewSource(s)
	driver.WithBind(Settings)

	// 绑定单个结构体数据
	//driver.WithBindKey("extend", Settings.Extend)

	// 初始化配置
	_cfg = &settings{
		settings:  *Settings,
		callbacks: fs,
	}
	_cfg.init()
}
