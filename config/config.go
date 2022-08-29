package config

import (
	"fmt"
	"github.com/cpyun/cpyun-admin-core/config/driver"
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
	//_cfg         config.Server
)

// @title    Setup
// @description   Setup 载入配置文件
// @auth      caillen             时间（2022/7/22   10:00 ）
// @param     s         string        "配置文件路径"
// @param     fs        func          "回调函数"
// @return
func Setup(s string, fs ...func()) {
	Settings = &Config{
		//Application: nil,
	}
	//1. 读取配置
	driver.NewSource(s)
	driver.WithEntity(Settings)

	// 日志
	Settings.Logger.Setup()
	_, err := Settings.Cache.Setup()
	if err != nil {
		fmt.Printf("=>>>>>>>>>>>>>>>>>> %s", err.Error())
	}

	fmt.Printf("%+v", Settings.Redis)
}
