package broadcasts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/impack-pratama/qontak/pkg"
	errs "github.com/impack-pratama/qontak/pkg/errors"
	"github.com/impack-pratama/qontak/whatsapp/broadcasts/direct_message"
	"net/http"
	"strings"
)

const (
	URI_SEND_DIRECT_MESSAGE = "/broadcasts/whatsapp/direct"
)

type service struct {
	client  pkg.Client
	token   string
	baseUrl string
}

func (s *service) SendDirectMessage(request *direct_message.SendDirectMessageRequest) (response *direct_message.SendDirectMessageResponse, err error) {
	var resp *http.Response
	var uri string
	var errorResponse errs.DefaultErrorResponse
	var r direct_message.SendDirectMessageResponse

	uri = fmt.Sprint(s.baseUrl, URI_SEND_DIRECT_MESSAGE)
	if resp, err = s.client.Execute(http.MethodPost, s.token, uri, request.ToJSON(), pkg.CONTENT_TYPE_JSON); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		return nil, errors.New(strings.Join(errorResponse.Error.Messages, ", "))
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return &r, err
}

func NewService(client pkg.Client, token string, baseUrl string) Service {
	a := new(service)
	a.client = client
	a.token = fmt.Sprint("Bearer ", token)
	a.baseUrl = baseUrl
	return a
}
