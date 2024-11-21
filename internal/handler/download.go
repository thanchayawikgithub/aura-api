package handler

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (s *AttachmentService) DownloadFile(ctx context.Context, path string) error {
	_, err := s.MinioClient.DownloadFile(ctx, bucketName, path, minio.GetObjectOptions{})
	return err
}
