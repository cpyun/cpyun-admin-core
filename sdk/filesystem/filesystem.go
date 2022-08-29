package filesystem

import (
	"context"
	"github.com/cpyun/cpyun-admin-core/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"path/filepath"
)

type FileSystem struct {
	ctx        context.Context
	client     *minio.Client
	bucketName string
}

//
func New() *FileSystem {
	ctx := context.Background()
	option := config.Settings.Storage.Minio

	s3Client, err := minio.New(option.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(option.AccessKeyID, option.SecretAccessKey, ""),
		Secure: option.Secure,
	})

	if err != nil {
		//log.Fatalln(err)
		log.Println(err)
	}

	return &FileSystem{
		client:     s3Client,
		ctx:        ctx,
		bucketName: option.Bucket,
	}
}

// 设置上下文
func (f *FileSystem) SetContext(ctx context.Context) *FileSystem {
	f.ctx = ctx
	return f
}

// 设置bucket
func (f *FileSystem) SetBucket(bucketName string) *FileSystem {
	f.bucketName = bucketName
	return f
}

// 保存文件
func (f FileSystem) PutFile(path string, file *multipart.FileHeader, rule string) (minio.UploadInfo, error) {
	name := generateHashName(rule) + filepath.Ext(file.Filename)
	return f.PutFileAs(path, file, name)
}

// 指定文件名保存文件
func (f FileSystem) PutFileAs(path string, file *multipart.FileHeader, name string) (minio.UploadInfo, error) {
	src, err := file.Open()
	if err != nil {
		log.Fatalln(err)
		return minio.UploadInfo{}, err
	}
	defer src.Close()

	putObjectOptions := minio.PutObjectOptions{
		ContentType: file.Header.Get(" Content-Type"),
	}

	objectName := path + `/` + name

	info, err := f.client.PutObject(f.ctx, f.bucketName, objectName, src, file.Size, putObjectOptions)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	return info, nil

	//fmt.Printf("uploadinfo: %+v \r\n", info)
	//{Bucket:notice Key:tmp/DIR_878_FW120B05_decode.BIN ETag:07d4fdc93ad1d270c06e1c924ee26b83 Size:11188937 LastModified:0001-01-01 00:00:00 +0000 UTC Location: VersionID: Expiration:0001-01-01 00:00:00 +0000 UTC ExpirationRuleID:}

}

func (f FileSystem) GetObject(objectName string) *minio.Object {
	getObjectOptions := minio.GetObjectOptions{}

	object, err := f.client.GetObject(f.ctx, f.bucketName, objectName, getObjectOptions)
	if err != nil {
		log.Fatalln(err)
		return &minio.Object{}
	}

	return object
}
