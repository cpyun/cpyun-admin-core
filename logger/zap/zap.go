package zap

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/cpyun/cpyun-admin-core/logger"
)

type zapLog struct {
	cfg  zap.Config
	zap  *zap.Logger
	opts logger.Options
	sync.RWMutex
	fields map[string]interface{}
}

func (l *zapLog) Init(opts ...logger.Option) error {
	//var err error

	for _, o := range opts {
		o(&l.opts)
	}

	zapConfig := zap.NewProductionConfig()
	if zConfig, ok := l.opts.Context.Value(configKey{}).(zap.Config); ok {
		zapConfig = zConfig
	}

	if ecConfig, ok := l.opts.Context.Value(encoderConfigKey{}).(zapcore.EncoderConfig); ok {
		zapConfig.EncoderConfig = ecConfig
	}

	writer, ok := l.opts.Context.Value(writerKey{}).(io.Writer)
	if !ok {
		writer = os.Stdout
	}

	skip, ok := l.opts.Context.Value(callerSkipKey{}).(int)
	if !ok || skip < 1 {
		skip = 1
	}

	// Set log Level if not default
	zapConfig.Level = zap.NewAtomicLevel()
	if l.opts.Level != logger.InfoLevel {
		zapConfig.Level.SetLevel(loggerToZapLevel(l.opts.Level))
	}

	// 设置日志时间格式
	if timeFormat, ok := l.opts.Context.Value(timeFormatKey{}).(string); ok {
		if timeFormat != "" {
			zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)
		} else {
			zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		}
	}

	// 日志级别大写
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer)),
		zapConfig.Level)

	log := zap.New(logCore, zap.AddCaller(), zap.AddCallerSkip(skip), zap.AddStacktrace(zap.DPanicLevel))
	//log, err := zapConfig.Build(zap.AddCallerSkip(skip))
	//if err != nil {
	//	return err
	//}

	// Adding seed fields if exist
	if l.opts.Fields != nil {
		data := []zap.Field{}
		for k, v := range l.opts.Fields {
			data = append(data, zap.Any(k, v))
		}
		log = log.With(data...)
	}

	// Adding namespace
	if namespace, ok := l.opts.Context.Value(namespaceKey{}).(string); ok {
		log = log.With(zap.Namespace(namespace))
	}

	// defer log.Sync() ??

	l.cfg = zapConfig
	l.zap = log
	l.fields = make(map[string]interface{})

	return nil
}

func (l *zapLog) Fields(fields map[string]interface{}) logger.Logger {
	l.Lock()
	nfields := make(map[string]interface{}, len(l.fields))
	for k, v := range l.fields {
		nfields[k] = v
	}
	l.Unlock()
	for k, v := range fields {
		nfields[k] = v
	}

	data := make([]zap.Field, 0, len(nfields))
	for k, v := range fields {
		data = append(data, zap.Any(k, v))
	}

	zl := &zapLog{
		cfg:    l.cfg,
		zap:    l.zap,
		opts:   l.opts,
		fields: nfields,
	}

	return zl
}

func (l *zapLog) Error(err error) logger.Logger {
	return l.Fields(map[string]interface{}{"error": err})
}

func (l *zapLog) Log(level logger.Level, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zapLog) Logf(level logger.Level, format string, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zapLog) String() string {
	return "zap"
}

func (l *zapLog) Options() logger.Options {
	return l.opts
}

//
func NewZap(opts ...logger.Option) (logger.Logger, error) {
	// Default options
	options := logger.Options{
		Level:   logger.InfoLevel,
		Fields:  make(map[string]interface{}),
		Out:     os.Stderr,
		Context: context.Background(),
	}

	l := &zapLog{opts: options}
	if err := l.Init(opts...); err != nil {
		return nil, err
	}

	return l, nil
}

//
func loggerToZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return zap.DebugLevel
	case logger.InfoLevel:
		return zap.InfoLevel
	case logger.WarnLevel:
		return zap.WarnLevel
	case logger.ErrorLevel:
		return zap.ErrorLevel
	case logger.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func zapToLoggerLevel(level zapcore.Level) logger.Level {
	switch level {
	case zap.DebugLevel:
		return logger.DebugLevel
	case zap.InfoLevel:
		return logger.InfoLevel
	case zap.WarnLevel:
		return logger.WarnLevel
	case zap.ErrorLevel:
		return logger.ErrorLevel
	case zap.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
