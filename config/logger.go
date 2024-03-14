package config

import "github.com/cpyun/gyopls-core/sdk/pkg/logger"

var LoggerConfig = new(Logger)

type Logger struct {
	Type          string    `mapstructure:"type" json:"type" yaml:"type"`
	Level         string    `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	TimeFormat    string    `mapstructure:"time-format" json:"time-format" yaml:"time-format"`          // 时间格式
	Prefix        string    `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Path          string    `mapstructure:"path" json:"path"  yaml:"path"`                              // 日志文件
	ShowLine      bool      `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	EncodeLevel   string    `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string    `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	Stdout        string    `mapstructure:"stdout" json:"stdout" yaml:"stdout"`                         // 输出控制台
	Cut           LoggerCut `mapstructure:"cut" json:"cut" yaml:"cut"`                                  // 日志裁切
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithTimeFormat(e.TimeFormat),
		// 日志裁切
		logger.WithCap(e.Cut.Cap),
		logger.WithMaxAge(e.Cut.MaxAge),
		logger.WithMaxBackups(e.Cut.MaxBackups),
		logger.WithCompress(e.Cut.Compress),
	)
}

type LoggerCut struct {
	Cap        uint `mapstructure:"cap" json:"cap" yaml:"cap"`                     // 裁切 MB
	MaxAge     uint `mapstructure:"max-age" json:"max-age" yaml:"max-age"`         // 最大备份天数
	MaxBackups uint `mapstructure:"max-age" json:"max-backups" yaml:"max-backups"` // 最大备份数
	Compress   bool `mapstructure:"compress" json:"compress" yaml:"compress"`      // 是否压缩
}
