package logger

import (
	"io"
	"os"

	log "github.com/cpyun/cpyun-admin-core/logger"
	"github.com/cpyun/cpyun-admin-core/logger/zap"

	"gopkg.in/natefinch/lumberjack.v2"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...OptionFunc) log.Logger {
	var err error

	op := setDefault()
	for _, o := range opts {
		o(&op)
	}

	var output io.Writer
	switch op.stdout {
	case "file":
		output = &lumberjack.Logger{
			Filename:   op.path,            // 文件名
			MaxSize:    int(op.cap),        // 单位：MB
			MaxAge:     int(op.maxAge),     // 最大备份天数
			MaxBackups: int(op.maxBackups), // 最大备份数
			Compress:   op.compress,        // 是否压缩
			LocalTime:  true,               // 本地时间
		}
	default:
		output = os.Stdout
	}

	var level log.Level
	level, err = log.GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}

	//
	switch op.driver {
	case "zap":
		log.DefaultLogger, err = zap.NewZap(
			log.WithLevel(level),
			zap.WithTimeFormat(op.timeFormat),
			zap.WithOutput(output),
			zap.WithCallerSkip(2),
		)
		if err != nil {
			log.Fatalf("new zap logger error, %s", err.Error())
		}
	case "logrus":
		log.Fatal("not support logrus")
	default:
		log.DefaultLogger = log.NewLogger(
			log.WithLevel(level),
			log.WithOutput(output),
		)
	}

	return log.DefaultLogger
}
