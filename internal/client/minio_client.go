package client

import (
	"aura/internal/config"
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type (
	IMinIOClient interface {
		UploadFile(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	}

	MinIOClient struct {
		*minio.Client
	}
)

func NewMinioClient(cfg *config.MinIO) *MinIOClient {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		log.Printf("error: minio client %#v\n", err)
	}

	return &MinIOClient{minioClient}
}

func (c *MinIOClient) UploadFile(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	return c.Client.PutObject(ctx, bucketName, objectName, reader, objectSize, opts)
}
