package driver

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

func NewSource(s string) string {
	viper.SetConfigFile(s)

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

// WithBind 绑定数据 alias别名 > 调用Set设置 > flag > env > config > key/value store > default
func WithBind(e any) {
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
		fmt.Printf("[config] ============> with bind key [%s] error : %s\n", key, err.Error())
		panic(err)
	}
}

func WithBindEnv(prefix string) {
	// 获取环境变量
	viper.SetEnvPrefix(prefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	//envBindings := []struct{ key, envName string }{
	//	{key: "app_name", envName: "APP_NAME"},
	//	{key: "app_env", envName: "APP_ENV"},
	//	{key: "app_debug", envName: "APP_DEBUG"},
	//	{key: "app_url", envName: "APP_URL"},
	//	{key: "app_timezone", envName: "APP_TIMEZONE"},
	//}
	//for _, binding := range envBindings {
	//	err := viper.BindEnv(binding.key, binding.envName)
	//	if err != nil {
	//		fmt.Println("[config] ============> with bind env error :", err.Error())
	//	}
	//}

}
