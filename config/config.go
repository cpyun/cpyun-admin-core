package config

import "log"

type Config struct {
	Application *Application          `mapstructure:"application" json:"application" yaml:"application"`
	Database    *Database             `json:"database" yaml:"database"`
	Databases   *map[string]*Database `json:"databases" yaml:"databases"`
	Logger      *Logger               `mapstructure:"logger" json:"logger" yaml:"logger"`
	Cache       *Cache                `mapstructure:"cache" yaml:"cache" json:"cache"`
	Filesystem  *Filesystem           `mapstructure:"filesystem" json:"mysql" yaml:"filesystem"`
	Queue       *Queue                `json:"queue" yaml:"queue"`
	Locker      *Locker               `json:"locker" yaml:"locker"`
	JWT         *Jwt                  `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis       Redis                 `mapstructure:"redis" json:"redis" yaml:"redis"`
	Casbin      *Casbin               `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Extend      interface{}           `yaml:"extend"`
	//
	opts setOptions
}

func (c *Config) OnChange() {
	//c.init()
	log.Println("config change and reload")
}

//多db改造
func (c *Config) multiDatabase() {
	if len(*c.Databases) == 0 {
		*c.Databases = map[string]*Database{
			"*": c.Database,
		}
	}
}

func (c *Config) runCallback() {
	for _, callback := range c.opts.callbacks {
		callback()
	}
}

func (c *Config) init() {
	c.Logger.Setup()
	c.multiDatabase()

	//调用回调函数
	c.runCallback()
}
