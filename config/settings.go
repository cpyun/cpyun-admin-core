package config

import (
	"github.com/cpyun/gyopls-core/config/driver"
	"github.com/cpyun/gyopls-core/config/driver/source"
	"log"
)

var (
	ExtendConfig interface{}
	Settings     *Config
	//_cfg         *settings
)

//type settings struct {
//	prefix    string // 前缀
//	settings  Config
//	callbacks []func()
//}
//
//// 初始化
//func (e *settings) init() {
//	// 配置日志
//	e.settings.Logger.Setup()
//	//配置多数据库
//	e.settings.multiDatabase()
//
//	// 调用回调函数
//	e.runCallback()
//}
//
//// 回调函数
//func (e *settings) runCallback() {
//	for _, callback := range e.callbacks {
//		callback()
//	}
//}
//
//// OnChange 修改配置
//func (e *settings) OnChange() {
//	e.init()
//	log.Println("!!! config change and reload")
//}

// Setup
// @description   Setup 载入配置文件
// @auth      cpYun             时间（2022/7/22   10:00 ）
// @param     s         string        "配置文件路径"
// @param     fs        func          "回调函数"
// @return
func Setup(fs ...SetOptionFuc) {
	var err error

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

	for _, f := range fs {
		f(&Settings.opts)
	}

	//
	driver.DefaultConfig, err = driver.NewConfig(
		driver.WithSource(Settings.opts.source...),
		driver.WithEntity(Settings),
	)
	if err != nil {
		log.Fatalln("new config object fail: ", err.Error())
	}

	Settings.init()

	// 初始化配置
	//_cfg = &settings{
	//	settings:  *Settings,
	//	callbacks: fs,
	//}
	//_cfg.init()
}

type setOptions struct {
	source    []source.Source
	callbacks []func()
}

type SetOptionFuc func(*setOptions)

func WithSource(s ...source.Source) SetOptionFuc {
	return func(o *setOptions) {
		o.source = append(o.source, s...)
	}

}

func WithCallback(fs ...func()) SetOptionFuc {
	return func(o *setOptions) {
		o.callbacks = append(o.callbacks, fs...)
	}
}
