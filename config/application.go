package config

type Application struct {
	Mode         string `json:"mode" yaml:"mode"`                   //环境配置 dev开发环境 test测试环境 prod线上环境
	Host         string `json:"host" yaml:"host"`                   // 服务器ip，默认使用 0.0.0.0
	Port         int    `json:"port" yaml:"port"`                   // 端口号
	Name         string `json:"name" yaml:"name"`                   // 服务名称
	ReadTimeout  int    `json:"read-timeout" yaml:"read-timeout"`   // 读超时
	WriteTimeout int    `json:"write-timeout" yaml:"write-timeout"` // 写超时
	EnableDp     bool   `json:"enable-dp" yaml:"enable-dp"`         // 数据权限功能开关
}

var ApplicationConfig = new(Application)
