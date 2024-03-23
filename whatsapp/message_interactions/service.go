package message_interactions

import (
	"encoding/json"
	"fmt"
	"github.com/impack-pratama/qontak/pkg"
	"net/http"
)

const (
	QONTAK_MESSAGE_INTERACTION_URI = "/message_interactions"
)

type service struct {
	client  pkg.Client
	token   string
	baseUrl string
}

type EnableDisableMessageInteractions struct {
	StatusTemplate bool   `json:"status_template"`
	Url            string `json:"url"`
}

func (e *EnableDisableMessageInteractions) ToJSON() []byte {
	b, _ := json.Marshal(e)
	return b
}

func (s *service) EnableMessageInteractions(request *EnableMessageInteractionsRequest) (err error) {
	payload := new(EnableDisableMessageInteractions)
	payload.StatusTemplate = true
	payload.Url = request.Url

	return s.enableOrDisableMessageInteraction(payload)
}

func (s *service) DisableMessageInteractions(request *DisableMessageInteractionsRequest) (err error) {
	payload := new(EnableDisableMessageInteractions)
	payload.StatusTemplate = true
	payload.Url = request.Url

	return s.enableOrDisableMessageInteraction(payload)
}

func (s *service) enableOrDisableMessageInteraction(request *EnableDisableMessageInteractions) (err error) {
	var resp *http.Response

	uri := fmt.Sprint(s.baseUrl, QONTAK_MESSAGE_INTERACTION_URI)
	if resp, err = s.client.Execute(http.MethodPut, s.token, uri, request.ToJSON(), pkg.CONTENT_TYPE_JSON); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to enable message interactions")
	}

	return nil
}

func NewService(client pkg.Client, token string, baseUrl string) Service {
	a := new(service)
	a.client = client
	a.token = fmt.Sprint("Bearer ", token)
	a.baseUrl = baseUrl
	return a
}
