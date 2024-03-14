package zap

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/cpyun/gyopls-core/logger"
)

type Options struct {
	logger.Options
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetContext(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetContext(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetContext(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) logger.Option {
	return logger.SetContext(namespaceKey{}, namespace)
}

type writerKey struct{}

func WithOutput(out io.Writer) logger.Option {
	return logger.SetContext(writerKey{}, out)
}

// 时间格式
type timeFormatKey struct{}

func WithTimeFormat(timeFormat string) logger.Option {
	return logger.SetContext(timeFormatKey{}, timeFormat)
}
