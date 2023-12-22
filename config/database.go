package config

var (
	DatabaseConfig  = new(Database)
	DatabasesConfig = make(map[string]*Database)
)

type Database struct {
	Driver          string             `mapstructure:"driver" json:"driver" yaml:"driver"`
	Source          string             `mapstructure:"source" json:"source" yaml:"source"`
	ConnMaxIdleTime int                `mapstructure:"conn-max-idle-time" json:"conn-max-idle-time" yaml:"conn-max-idle-time"`
	ConnMaxLifeTime int                `mapstructure:"conn-max-life-time" json:"conn-max-life-time" yaml:"conn-max-life-time"`
	MaxIdleConns    int                `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns    int                `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LoggerMode      string             `mapstructure:"logger-mode" json:"logger-mode" yaml:"logger-mode"`
	Registers       []DBResolverConfig `mapstructure:"registers" json:"registers" yaml:"registers"`
}

type DBResolverConfig struct {
	Sources  []string `mapstructure:"sources" json:"sources" yaml:"sources"`
	Replicas []string `mapstructure:"replicas" json:"replicas" yaml:"replicas"`
	Policy   string   `mapstructure:"policy" json:"policy" yaml:"policy"`
	Tables   []string `mapstructure:"tables" json:"tables" yaml:"tables"`
}
