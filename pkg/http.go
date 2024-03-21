package pkg

import (
	"bytes"
	"net/http"
	"time"
)

type Client interface {
	Execute(method string, token string, uri string, body []byte, contentType string) (response *http.Response, err error)
	GetClient() *http.Client
}

type httpClient struct {
	client *http.Client
}

func (h *httpClient) GetClient() *http.Client {
	return h.client
}

func (h *httpClient) Execute(method string, token string, uri string, body []byte, contentType string) (response *http.Response, err error) {
	var req *http.Request

	if req, err = http.NewRequest(method, uri, bytes.NewReader(body)); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", token)

	return h.client.Do(req)
}

func NewHttp() Client {
	return &httpClient{client: &http.Client{Timeout: 10 * time.Second}}
}
