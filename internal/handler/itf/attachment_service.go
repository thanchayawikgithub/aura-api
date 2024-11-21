package itf

import (
	"context"
	"mime/multipart"
)

type IAttachmentService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) error
	DownloadFile(ctx context.Context, path string) error
}
