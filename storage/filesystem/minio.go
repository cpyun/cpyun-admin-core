package filesystem

import (
	"context"
	"errors"
	"github.com/cpyun/cpyun-admin-core/storage"
	"mime/multipart"
	"path/filepath"
	"reflect"

	"github.com/minio/minio-go/v7"

	log "github.com/cpyun/cpyun-admin-core/logger"
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
		log.Fatal("File system initialization failed, err: ", err)
		return nil
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
func (m *Minio) SetBucket(bucketName string) {
	m.bucketName = bucketName
}

// PutFile 保存文件
// rootPath 文件根目录 (eg: “”|“words/”|“images/”)
// file 文件
// rule string 如果为空，则由系统自动生成（带目录,不带后缀）(eg:"20200101/file")
func (m *Minio) PutFile(rootPath string, file *multipart.FileHeader, rule string) (minio.UploadInfo, error) {
	name := generateHashName(rule) + filepath.Ext(file.Filename)
	filePath := rootPath + name
	return m.PutFileAs(filePath, file)
}

// PutFileAs 指定文件名保存文件
func (m *Minio) PutFileAs(path string, file *multipart.FileHeader) (minio.UploadInfo, error) {
	src, err := file.Open()
	if err != nil {
		log.Error("File save failed, err: ", err)
		return minio.UploadInfo{}, err
	}
	defer src.Close()

	putObjectOptions := minio.PutObjectOptions{
		ContentType: file.Header.Get(" Content-Type"),
	}

	objectName := path

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
		log.Error("Failed to obtain file, err: ", err)
		return &minio.Object{}
	}
	defer object.Close()

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

func (m Minio) GetStore() {

}

// GetClient 获取客户端
func (m *Minio) GetClient() *minio.Client {
	return m.client
}

func CovertInterfaceToStruct(face storage.AdapterFilesystem) (*Minio, error) {
	value := reflect.ValueOf(face)
	if value.IsNil() {
		return nil, errors.New("value is nil")
	} else if value.Kind() != reflect.Ptr {
		return nil, errors.New("error of kind [pointer]")
	}

	//// 取数据
	//value = value.Elem()
	//if value.Kind() != reflect.Struct {
	//	return minio, errors.New("not a struct")
	//}

	return value.Interface().(*Minio), nil
}
