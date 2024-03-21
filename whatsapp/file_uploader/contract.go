package file_uploader

import "context"

type Service interface {
	// UploadFromUrl Upload a file to Qontak server, if there is any document you want to send via qontak whatsapp
	UploadFromUrl(ctx context.Context, filename string, url string) (response *UploadResponse, err error)
}
