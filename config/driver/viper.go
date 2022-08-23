package driver

import (
	"context"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type entity interface {
	//OnChange()
}

type Options struct {
	entity entity
}

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

func WithEntity(e entity) {
	_, cancel := context.WithCancel(context.Background())
	// 绑定数据
	if err := viper.Unmarshal(e); err != nil {
		panic(err)
	}

	// 监听和重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(et fsnotify.Event) {
		fmt.Println("[config] ============> config file changed:", et.Name)
		if err := viper.Unmarshal(e); err != nil {
			fmt.Println(err)
		}
		cancel()
	})
}
