package file_uploader

import "context"

type Service interface {
	// UploadFromUrl Upload a file to Qontak server, if there is any document you want to send via qontak whatsapp
	//if there are no content type, the default content type will be application/octet-stream
	UploadFromUrl(ctx context.Context, filename string, url string) (response *UploadResponse, err error)
	//UploadFromUrlWithContentType Upload a file with mime type application/octet-stream to Qontak server
	UploadFromUrlWithContentType(ctx context.Context, filename string, url string, contentType string) (response *UploadResponse, err error)
}
