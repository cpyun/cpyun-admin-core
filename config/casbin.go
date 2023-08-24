package config

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"model-path" yaml:"model-path"` //存放casbin模型的相对路径
	Adapter   string `mapstructure:"adapter" json:"adapter" yaml:"adapter"`          //适配器
	Watcher   string `mapstructure:"watcher" json:"watcher" yaml:"watcher"`          //监听
	Redis     *Redis
}

var CasbinConfig = new(Casbin)
