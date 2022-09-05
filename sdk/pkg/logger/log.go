package logger

import (
	"github.com/cpyun/cpyun-admin-core/logger/zap"
	"io"
	"os"

	"github.com/cpyun/cpyun-admin-core/debug/writer"
	//zap "github.com/cpyun/cpyun-admin-core/logger/driver"
	"github.com/cpyun/cpyun-admin-core/sdk/pkg"

	log "github.com/cpyun/cpyun-admin-core/logger"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...OptionFunc) log.Logger {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	if !pkg.PathExist(op.path) {
		err := pkg.PathCreate(op.path)
		if err != nil {
			log.Fatalf("create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch op.stdout {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPath(op.path),
			writer.WithCap(op.cap<<10),
		)
		if err != nil {
			log.Fatal("logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	var level log.Level
	level, err = log.GetLevel(op.level)
	if err != nil {
		log.Fatalf("get logger level error, %s", err.Error())
	}

	switch op.driver {
	case "zap":
		log.DefaultLogger, err = zap.NewLogger(log.WithLevel(level), log.WithOutput(output), zap.WithCallerSkip(2))
		if err != nil {
			log.Fatalf("new zap logger error, %s", err.Error())
		}
	//case "logrus":
	//	setLogger = logrus.NewLogger(logger.WithLevel(level), logger.WithOutput(output), logrus.ReportCaller())
	default:
		log.DefaultLogger = log.NewLogger(log.WithLevel(level), log.WithOutput(output))
	}
	return log.DefaultLogger
}
