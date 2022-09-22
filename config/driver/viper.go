package driver

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSource(s string) string {
	//1. 读取配置
	if s != "" {
		viper.SetConfigFile(s)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("settings")
	}

	// 获取环境变量
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	notFound := &viper.ConfigFileNotFoundError{}

	switch {
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		cobra.CheckErr(notFound.Error())
		break
	default:
		fmt.Fprintln(os.Stdout, "[config] ============> Using config file:", viper.ConfigFileUsed())
	}

	return ""
}

func WithBind(e any) {
	// 绑定数据
	if err := viper.Unmarshal(e); err != nil {
		panic(err)
	}

	// 监听和重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(et fsnotify.Event) {
		fmt.Println("[config] ============> config file changed:", et.Name)
		if err := viper.Unmarshal(e); err != nil {
			//opts.OnChange()
			fmt.Println(err)
		}
	})
}

//
func WithBindKey(key string, rawVal any) {
	err := viper.UnmarshalKey(key, rawVal)
	if err != nil {
		log.Fatal("[config] ============> error :", err.Error())
	}
}
