package runtime

import (
	"context"
	"github.com/cpyun/cpyun-admin-core/storage"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type FileSystem struct {
	ctx        context.Context
	store      storage.AdapterFilesystem
	bucketName string
}

func NewFilesystem(bucketName string, store storage.AdapterFilesystem) storage.AdapterFilesystem {
	ctx := context.Background()

	return &FileSystem{
		ctx:        ctx,
		store:      store,
		bucketName: bucketName,
	}
}

//
func (f *FileSystem) String() string {
	if f.store == nil {
		return ""
	}
	return f.store.String()
}

// PutFile 保存文件
func (f *FileSystem) PutFile(path string, file *multipart.FileHeader, rule string) (minio.UploadInfo, error) {
	return f.store.PutFile(path, file, rule)
}

// GetObject 获取
func (f *FileSystem) GetObject(objectName string) *minio.Object {
	return f.store.GetObject(objectName)
}

func (f *FileSystem) RemoveObject(file string) bool {
	return f.store.RemoveObject(file)
}

// SetContext 设置上下文
func (f *FileSystem) SetContext(ctx context.Context) *FileSystem {
	f.ctx = ctx
	return f
}

// SetBucket 设置bucket
func (f *FileSystem) SetBucket(bucketName string) *FileSystem {
	f.bucketName = bucketName
	return f
}

// GetStore ""
func (f *FileSystem) GetStore() storage.AdapterFilesystem {
	return f.store
}
