package config

import "github.com/cpyun/cpyun-admin-core/sdk/pkg/logger"

type Logger struct {
	Type          string `mapstructure:"type" json:"type" yaml:"type"`
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	TimeFormat    string `mapstructure:"time-format" json:"time-format" yaml:"time-format"`          // 时间格式
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Path          string `mapstructure:"path" json:"path"  yaml:"path"`                              // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	Stdout        string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`                         // 输出控制台
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithTimeFormat(e.TimeFormat),
		//logger.WithCap(e.Cap),
	)
}

var LoggerConfig = new(Logger)
