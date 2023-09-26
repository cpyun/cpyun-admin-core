package config

// System
// Deprecated: Use Application instead
type System struct {
	Env            string `mapstructure:"env" json:"env" yaml:"env"`                                     // 环境值
	Addr           int    `mapstructure:"addr" json:"addr" yaml:"addr"`                                  // 端口值
	DatabaseDriver string `mapstructure:"database-driver" json:"database-driver" yaml:"database-driver"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType        string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                      // Oss类型
	UseMultipoint  bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`    // 多点登录拦截
	UseRedis       bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	LimitCountIP   int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP    int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
