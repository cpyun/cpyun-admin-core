package storage

import (
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"time"

	"github.com/bsm/redislock"
)

const (
	PrefixKey = "__host"
)

type AdapterCache interface {
	String() string
	SetPrefix(string)
	Get(key string) (string, error)
	Set(key string, val interface{}, expire int) error
	Del(key string) error
	HashGet(hk, key string) (string, error)
	HashDel(hk, key string) error
	Increase(key string) error
	Decrease(key string) error
	Expire(key string, dur time.Duration) error
}

type AdapterQueue interface {
	String() string
	Append(message Messager) error
	Register(name string, f ConsumerFunc)
	Run()
	Shutdown()
}

type Messager interface {
	SetID(string)
	SetStream(string)
	SetValues(map[string]interface{})
	GetID() string
	GetStream() string
	GetValues() map[string]interface{}
	GetPrefix() string
	SetPrefix(string)
	SetErrorCount()
	GetErrorCount() int
}

type ConsumerFunc func(Messager) error

type AdapterLocker interface {
	String() string
	Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error)
}

// AdapterFilesystem 文件系统
type AdapterFilesystem interface {
	String() string
	PutFile(rootPath string, file *multipart.FileHeader, rule string) (minio.UploadInfo, error)
	GetObject(objectName string) *minio.Object
	RemoveObject(file string) bool
}
