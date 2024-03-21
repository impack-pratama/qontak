package file_uploader

type UploadResponseSuccess struct {
	FileName string `json:"filename"`
	Url      string `json:"url"`
}

type UploadResponse struct {
	Status string                `json:"status,omitempty"`
	Data   UploadResponseSuccess `json:"data"`
}
