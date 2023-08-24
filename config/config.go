package config

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
}

//多db改造
func (s *Config) multiDatabase() {
	if len(*s.Databases) == 0 {
		*s.Databases = map[string]*Database{
			"*": s.Database,
		}
	}
}
