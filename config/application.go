package config

type Application struct {
	Mode         string //环境配置 dev开发环境 test测试环境 prod线上环境
	Host         string // 服务器ip，默认使用 0.0.0.0
	Port         int    // 端口号
	Name         string // 服务名称
	ReadTimeOut  int
	WriteTimeOut int
	EnableDp     bool // 数据权限功能开关
}

var ApplicationConfig = new(Application)
