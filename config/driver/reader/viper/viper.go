package viper

import (
	"github.com/cpyun/gyopls-core/config/driver/reader"
	"github.com/cpyun/gyopls-core/config/driver/source"
	"github.com/spf13/viper"
)

type viperReader struct {
	opts Options
}

func (v *viperReader) Merge(sources ...*source.ChangeSet) (*source.ChangeSet, error) {
	return sources[0], nil

}

func (v *viperReader) Values(ch *source.ChangeSet) (reader.Values, error) {
	//if ch == nil {
	//	return nil, errors.New("change set is nil")
	//}

	return newValues(ch)
}

func (v *viperReader) String() string {
	return "viper"
}

func NewReaderViper() reader.Reader {
	return &viperReader{}
}

// viperValues Values
// interface
//
type viperValues struct {
	ch *source.ChangeSet
}

func (v *viperValues) Bytes() []byte {
	return v.ch.Data
}

func (v *viperValues) Get(path ...string) reader.Value {
	return nil
}

func (v *viperValues) Set(val interface{}, path ...string) {}

func (v *viperValues) Del(path ...string) {}

func (v *viperValues) Map() map[string]interface{} {
	return viper.AllSettings()
}

func (v *viperValues) Scan(e interface{}) error {
	return viper.Unmarshal(e)
}

func newValues(ch *source.ChangeSet) (reader.Values, error) {
	return &viperValues{ch: ch}, nil
}
