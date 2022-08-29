package config

type Storage struct {
	FilesystemCloud string
	Local           interface{}
	Minio           Minio
	Qiniu           Qiniu
	AliyunOSS       AliyunOSS
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key" json:"secret-access-key" yaml:"secret-access-key"`
	Secure          bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
	Region          string `mapstructure:"region" json:"region" yaml:"region"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
}

type Qiniu struct {
}

type AliyunOSS struct {
}
