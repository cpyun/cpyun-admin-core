package driver

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func NewSource(s string) string {
	viper.SetConfigFile(s)
	// 获取环境变量
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	notFound := &viper.ConfigFileNotFoundError{}
	if errors.As(err, notFound) {
		fmt.Println("[config] ============> config file not found: ", s)
		panic(notFound)
	} else if err != nil {
		fmt.Println("[config] ============> read in config error: ", err)
		panic(err)
	}

	fmt.Println("[config] ============> Using config file:", viper.ConfigFileUsed())
	return s
}

func WithBind(e any) {
	// 绑定数据
	if err := viper.Unmarshal(e); err != nil {
		fmt.Println("[config] ============> with bind parameter error: ", err)
		panic(err)
	}

	// 监听和重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(et fsnotify.Event) {
		fmt.Println("[config] ============> config file changed:", et.Name)
		if err := viper.Unmarshal(e); err != nil {
			fmt.Println("[config] ===========> change config and bind parameter error: ", err)
		}
	})
}

// WithBindKey 绑定单个
func WithBindKey(key string, rawVal any) {
	err := viper.UnmarshalKey(key, rawVal)
	if err != nil {
		fmt.Println("[config] ============> with bind key error :", err.Error())
	}
}
