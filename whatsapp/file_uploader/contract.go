package file_uploader

import "context"

type S3Config struct {
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
	Key       string
}

type R2Config struct {
	AccessKey string
	SecretKey string
	AccountID string
	Bucket    string
	Key       string
}

type Service interface {
	// UploadFromUrl Upload a file to Qontak server, if there is any document you want to send via qontak whatsapp
	//if there are no content type, the default content type will be application/octet-stream
	UploadFromUrl(ctx context.Context, filename string, url string) (response *UploadResponse, err error)
	//UploadFromUrlWithContentType Upload a file with mime type application/octet-stream to Qontak server
	UploadFromUrlWithContentType(ctx context.Context, filename string, url string, contentType string) (response *UploadResponse, err error)
	// UploadFromS3 Upload a file from AWS S3 to Qontak server
	UploadFromS3(ctx context.Context, filename string, config S3Config) (response *UploadResponse, err error)
	// UploadFromCloudflareR2 Upload a file from Cloudflare R2 to Qontak server
	UploadFromCloudflareR2(ctx context.Context, filename string, config R2Config) (response *UploadResponse, err error)
}
