package handler

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

const (
	bucketName = "int202411003"
)

func (s *AttachmentService) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = s.MinioClient.UploadFile(ctx, bucketName, file.Filename, f, file.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
