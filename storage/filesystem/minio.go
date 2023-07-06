package filesystem

import (
	"context"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

type Minio struct {
	ctx        context.Context
	client     *minio.Client
	bucketName string
}

func NewMinio(endpoint string, options *minio.Options, bucketName string) *Minio {
	ctx := context.Background()

	s3Client, err := minio.New(endpoint, options)
	if err != nil {
		log.Println(err)
	}

	return &Minio{
		client:     s3Client,
		ctx:        ctx,
		bucketName: bucketName,
	}
}

//
func (m *Minio) String() string {
	return "minio"
}

// SetContext 设置上下文
func (m *Minio) SetContext(ctx context.Context) *Minio {
	m.ctx = ctx
	return m
}

// SetBucket 设置bucket
func (m *Minio) SetBucket(bucketName string) *Minio {
	m.bucketName = bucketName
	return m
}

// PutFile 保存文件
func (m *Minio) PutFile(path string, file *multipart.FileHeader, rule string) (minio.UploadInfo, error) {
	name := generateHashName(rule) + filepath.Ext(file.Filename)
	return m.PutFileAs(path, file, name)
}

// PutFileAs 指定文件名保存文件
func (m *Minio) PutFileAs(path string, file *multipart.FileHeader, name string) (minio.UploadInfo, error) {
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

	info, err := m.client.PutObject(m.ctx, m.bucketName, objectName, src, file.Size, putObjectOptions)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	return info, nil

	//fmt.Printf("uploadinfo: %+v \r\n", info)
	//{Bucket:notice Key:tmp/DIR_878_FW120B05_decode.BIN ETag:07d4fdc93ad1d270c06e1c924ee26b83 Size:11188937 LastModified:0001-01-01 00:00:00 +0000 UTC Location: VersionID: Expiration:0001-01-01 00:00:00 +0000 UTC ExpirationRuleID:}

}

// GetObject 获取文件详情
func (m *Minio) GetObject(objectName string) *minio.Object {
	getObjectOptions := minio.GetObjectOptions{}

	object, err := m.client.GetObject(m.ctx, m.bucketName, objectName, getObjectOptions)
	if err != nil {
		log.Fatalln(err)
		return &minio.Object{}
	}

	return object
}

// RemoveObject 删除文件
func (m *Minio) RemoveObject(objectName string) bool {
	err := m.client.RemoveObject(m.ctx, m.bucketName, objectName, minio.RemoveObjectOptions{
		GovernanceBypass: true,
	})
	if err != nil {
		return false
	}
	return true
}

// GetClient 获取客户端
func (m *Minio) GetClient() *minio.Client {
	return m.client
}
