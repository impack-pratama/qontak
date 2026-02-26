package file_uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/impack-pratama/qontak/pkg"
	errs "github.com/impack-pratama/qontak/pkg/errors"
)

const (
	QONTAK_FILE_UPLOADER_URI = "/file_uploader"
	DEFAULT_CONTENT_TYPE     = "application/octet-stream"
	CONTENT_TYPE_PDF         = "application/pdf"
)

type service struct {
	client  pkg.Client
	token   string
	baseUrl string
}

func (s *service) UploadFromUrlWithContentType(ctx context.Context, filename string, url string, contentType string) (response *UploadResponse, err error) {
	var resp *http.Response

	if resp, err = http.Get(url); err != nil {
		ctx.Err()
		return nil, err
	}
	defer resp.Body.Close()

	if contentType == "" {
		contentType = resp.Header.Get("Content-Type")
		if contentType == "" {
			contentType = DEFAULT_CONTENT_TYPE
		}
	}

	return s.upload(ctx, filename, resp.Body, contentType)
}

func (s *service) upload(ctx context.Context, filename string, reader io.Reader, contentType string) (response *UploadResponse, err error) {
	var resp *http.Response
	var fw io.Writer
	var errorResponse errs.DefaultErrorResponse
	var r UploadResponse

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	qontakUrl := fmt.Sprint(s.baseUrl, QONTAK_FILE_UPLOADER_URI)

	if fw, err = s.CreateWriter(writer, "file", filename, contentType); err != nil {
		ctx.Err()
		return nil, err
	}
	if _, err = io.Copy(fw, reader); err != nil {
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

func (s *service) UploadFromUrl(ctx context.Context, filename string, url string) (response *UploadResponse, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		ctx.Err()
		return nil, err
	}

	return s.UploadFromUrlWithContentType(ctx, filename, url, resp.Header.Get("content-type"))
}

func (s *service) UploadFromS3(ctx context.Context, filename string, s3Config S3Config) (response *UploadResponse, err error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(s3Config.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3Config.AccessKey, s3Config.SecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)
	output, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s3Config.Bucket),
		Key:    aws.String(s3Config.Key),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	contentType := aws.ToString(output.ContentType)
	if contentType == "" {
		contentType = DEFAULT_CONTENT_TYPE
	}

	return s.upload(ctx, filename, output.Body, contentType)
}

func (s *service) UploadFromCloudflareR2(ctx context.Context, filename string, r2Config R2Config) (response *UploadResponse, err error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2Config.AccessKey, r2Config.SecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2Config.AccountID))
	})

	output, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r2Config.Bucket),
		Key:    aws.String(r2Config.Key),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	contentType := aws.ToString(output.ContentType)
	if contentType == "" {
		contentType = DEFAULT_CONTENT_TYPE
	}

	return s.upload(ctx, filename, output.Body, contentType)
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

func NewService(client pkg.Client, token string, baseUrl string) Service {
	a := new(service)
	a.client = client
	a.token = fmt.Sprint("Bearer ", token)
	a.baseUrl = baseUrl
	return a
}
