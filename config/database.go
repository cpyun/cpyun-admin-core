package config

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type Mysql struct {
	Hostname        string             `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	HostPort        string             `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Config          string             `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Dbname          string             `mapstructure:"database" json:"database" yaml:"database"` //数据库名
	Username        string             `mapstructure:"username" json:"username" yaml:"username"` //用户名
	Password        string             `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Charset         string             `mapstructure:"charset" json:"charset" yaml:"charset"`
	ConnMaxIdleTime int                `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time"`
	ConnMaxLifeTime int                `mapstructure:"conn_max_life_time" json:"conn_max_life_time"`
	MaxIdleConns    int                `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns    int                `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode         string             `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap          string             `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
	Registers       []DBResolverConfig `mapstructure:"registers"`
}

type DBResolverConfig struct {
	Sources  []string `mapstructure:"sources" json:"sources" yaml:"sources"`
	Replicas []string `mapstructure:"replicas" json:"logZap" yaml:"replicas"`
	Policy   string   `mapstructure:"policy" json:"policy" yaml:"policy"`
	Tables   []string `mapstructure:"tables" json:"tables" yaml:"tables"`
}

// @Title	parseDsn
// @Description 解析pdo连接的dsn信息
// @param  config 	string  连接信息
// @return string
func (m *Mysql) GetMysqlDsn() string {
	// "root:123456@tcp(127.0.0.1:3306)/tsgz?charset=utf8mb4&parseTime=True&loc=Local"

	params := make(map[string]string)
	params["charset"] = m.Charset

	cfg := mysql.Config{
		User:             m.Username,
		Passwd:           m.Password,
		Net:              "tcp",
		Addr:             m.Hostname,
		DBName:           m.Dbname,
		Params:           params,
		Loc:              time.Local,
		Collation:        m.Charset + "_unicode_ci",
		MaxAllowedPacket: 4 << 20,

		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		ParseTime:            true,
	}
	dsn := cfg.FormatDSN()
	return dsn
}

// 获取日志模式
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
