package upload

import (
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/antdate/antdate-admin/internal/admin/curd"
	"mime/multipart"
	"path/filepath"
)

const (
	UploaderOSS = "oss"
)

type OSSUploader struct {
	client *oss.Client
	bucket string
}

func NewOSSUploader(endpoint, accessKeyID, accessKeySecret string, bucket string) (file.Uploader, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return OSSUploader{client: client, bucket: bucket}, nil
}

func (q OSSUploader) Upload(form *multipart.Form) error {
	// 接收一个表单类型，这里实现上传逻辑
	// 这里调用框架的辅助函数
	return file.Upload(func(f *multipart.FileHeader, name string) (string, error) {
		// 这里实现上传逻辑，返回文件路径与错误信息
		bucket, err := q.client.Bucket(q.bucket)
		if err != nil {
			return "", err
		}

		uploadFile, err := f.Open()
		if err != nil {
			return "", err
		}

		filename := modules.Uuid() + filepath.Ext(f.Filename)

		err = bucket.PutObject(filename, uploadFile)
		if err != nil {
			return "", err
		}
		return curd.GetOSSDomain(filename), nil
	}, form)

}
