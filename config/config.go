package config

import (
	"github.com/cpyun/cpyun-admin-core/config/driver"
	"log"
)

type Config struct {
	//System  		System  		`mapstructure:"system" json:"system" yaml:"system"`
	Application Application `mapstructure:"application" json:"application" yaml:"application"`
	JWT         Jwt         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//Zap     		Zap     		`mapstructure:"zap" json:"zap" yaml:"zap"`
	Logger Logger `mapstructure:"logger" json:"logger" yaml:"logger"`
	Cache  Cache  `mapstructure:"cache" yaml:"cache" json:"cache"`

	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Email   		Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`

	//Captcha 		Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	//AutoCode 		Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Databases *map[string]*Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Pgsql  		Pgsql 	`mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	//DBList 		[]DB 	`mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	//// oss
	Storage Storage `mapstructure:"storage" json:"mysql" yaml:"storage"`

	//
	//Excel Excel 	`mapstructure:"excel" json:"excel" yaml:"excel"`
	//Timer Timer 	`mapstructure:"timer" json:"timer" yaml:"timer"`
	//
	//// 跨域配置
	//Cors CORS 	`mapstructure:"cors" json:"cors" yaml:"cors"`
	Extend interface{} `yaml:"extend"`
}

// 多db改造
//func (e *Config) multiDatabase() {
//	if len(*e.Databases) == 0 {
//		*e.Databases = map[string]*Mysql{
//			//"*": e.Mysql,
//		}
//
//	}
//}

var (
	ExtendConfig interface{}
	Settings     *Config
	_cfg         *settings
)

type settings struct {
	Settings  Config
	callbacks []func()
}

// 初始化
func (e *settings) init() {
	e.Settings.Logger.Setup()
	_, _ = e.Settings.Cache.Setup()

	// 调用回调函数
	e.runCallback()
}

// 回调函数
func (e *settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

// 修改配置
func (e *settings) OnChange() {
	e.init()
	log.Println("!!! config change and reload")
}

// @title    Setup
// @description   Setup 载入配置文件
// @auth      caillen             时间（2022/7/22   10:00 ）
// @param     s         string        "配置文件路径"
// @param     fs        func          "回调函数"
// @return
func Setup(s string, fs ...func()) {
	//var err error

	Settings = &Config{
		Extend: ExtendConfig,
	}
	_cfg = &settings{
		Settings:  *Settings,
		callbacks: fs,
	}

	//1. 读取配置
	driver.NewSource(s)

	//driver.DefaultConfig, err = driver.NewConfig(
	//	// 绑定实体
	//	driver.WithEntity(_cfg),
	//)
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("New config object fail: %s", err.Error()))
	//}
	//entity := defaultConfig.Options().Entity
	//driver.WithBind(entity)
	//fmt.Printf("entity >>>>>>>>>>>>>>> %+v \n", entity)

	driver.WithBind(Settings)
	//fmt.Printf("Settings >>>>>>>>>>>>>>> %+v \n", Settings)
	_cfg.Settings = *Settings

	// 绑定单个结构体数据
	//driver.WithBindKey("application", ApplicationConfig)
	driver.WithBindKey("extend", Settings.Extend)

	// 初始化配置
	_cfg.init()
}
