package pkg

import (
	"bytes"
	"net/http"
	"time"
)

type Client interface {
	Execute(method string, token string, uri string, body []byte) (response *http.Response, err error)
}

type httpClient struct {
	client *http.Client
}

func (h *httpClient) Execute(method string, token string, uri string, body []byte) (response *http.Response, err error) {
	var req *http.Request

	if req, err = http.NewRequest(method, uri, bytes.NewReader(body)); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	return h.client.Do(req)
}

func NewHttp() Client {
	return &httpClient{client: &http.Client{Timeout: 10 * time.Second}}
}
