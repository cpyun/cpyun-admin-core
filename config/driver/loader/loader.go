package loader

import "github.com/cpyun/cpyun-admin-core/config/driver/source"

type Loader interface {
	Load(...source.Source) error
}
