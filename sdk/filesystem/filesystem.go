package filesystem

import (
	"context"
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
func New() FileSystem {
	ctx := context.Background()

	endpoint := "192.168.99.93:9090"
	accessKeyID := "LaBY0JlcsRH4pac7uM17wxCB"
	secretAccessKey := "neTSW#OrzWR7a^VrXvlp6QZ8bo!akWDMTy3b"
	useSSL := false
	bucketName := "notice"

	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return FileSystem{
		client:     s3Client,
		ctx:        ctx,
		bucketName: bucketName,
	}
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
