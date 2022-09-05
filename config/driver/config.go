package driver

import (
	"context"
)

type Config interface {
	// Options in the config
	Options() Options
}

//
type entity interface {
	OnChange()
}

//
type Options struct {
	Context context.Context

	Entity entity
}

type OptionFunc func(o *Options)

var (
	// DefaultConfig Default Config Manager
	DefaultConfig Config
)

// NewConfig returns new config
func NewConfig(opts ...OptionFunc) (Config, error) {
	return newConfig(opts...)
}
