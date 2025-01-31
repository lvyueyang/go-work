package utils

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type UploadOssOptions struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	File            io.Reader
	FilePath        string
	Options         []oss.Option
}

// UploadFileToAliOss 上传文件到阿里云oss
func UploadFileToAliOss(opt UploadOssOptions) error {
	client, err := oss.New(opt.Endpoint, opt.AccessKeyID, opt.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(opt.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(opt.FilePath, opt.File, opt.Options...)
	if err != nil {
		return err
	}
	return nil
}
