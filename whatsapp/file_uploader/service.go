package file_uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/impack-pratama/qontak/pkg"
	errs "github.com/impack-pratama/qontak/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

const (
	QONTAK_FILE_UPLOADER_URI = "/file_uploader"
)

type service struct {
	client  pkg.Client
	token   string
	baseUrl string
}

func (s *service) UploadFromUrl(ctx context.Context, filename string, url string) (response *UploadResponse, err error) {
	var resp *http.Response
	var fw io.Writer
	var errorResponse errs.DefaultErrorResponse
	var r UploadResponse

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	qontakUrl := fmt.Sprint(s.baseUrl, QONTAK_FILE_UPLOADER_URI)

	if resp, err = http.Get(url); err != nil {
		ctx.Err()
		return nil, err
	}
	defer resp.Body.Close()

	if fw, err = s.CreateWriter(writer, "file", filename, resp.Header.Get("content-type")); err != nil {
		ctx.Err()
		return nil, err
	}
	if _, err = io.Copy(fw, resp.Body); err != nil {
		ctx.Err()
		return nil, err
	}
	writer.Close()

	if resp, err = s.client.Execute(http.MethodPost, s.token, qontakUrl, body.Bytes(), writer.FormDataContentType()); err != nil {
		ctx.Err()
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		ctx.Err()
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		return nil, errors.New(strings.Join(errorResponse.Error.Messages, ", "))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return &r, err

}

func (s *service) CreateWriter(writer *multipart.Writer, fieldname, filename string, contentType string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			quoteEscaper.Replace(fieldname), quoteEscaper.Replace(filename)))
	h.Set("Content-Type", contentType)
	return writer.CreatePart(h)
}

func (s *service) UploadFromFile(ctx context.Context, file string) (response *UploadResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func NewService(client pkg.Client, token string, baseUrl string) Service {
	a := new(service)
	a.client = client
	a.token = fmt.Sprint("Bearer ", token)
	a.baseUrl = baseUrl
	return a
}
