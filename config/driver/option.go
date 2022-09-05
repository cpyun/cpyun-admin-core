package driver

//
func WithSource(s string) OptionFunc {
	return func(o *Options) {

	}
}

// 实体
func WithEntity(e entity) OptionFunc {
	return func(o *Options) {
		o.Entity = e
	}
}
