package driver

var (
	// DefaultConfig Default Config Manager
	DefaultConfig Config
)

type Config interface {
	// Options in the config
	Options() Options
}

//
type Entity interface {
	OnChange()
}

//
//type OptionFunc func(o *Options)

// NewConfig returns new config
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}
