package waba_interactions

import (
	"encoding/json"
	"fmt"
	"github.com/impack-pratama/qontak/pkg"
	"net/http"
)

const (
	QONTAK_WABA_INTERACTION_URI = "/waba_interactions"
)

type service struct {
	client  pkg.Client
	token   string
	baseUrl string
}

type EnableDisableWabaInteractions struct {
	StatusTemplate bool   `json:"status_template"`
	Url            string `json:"url"`
}

func (e *EnableDisableWabaInteractions) ToJSON() []byte {
	b, _ := json.Marshal(e)
	return b
}

func (s *service) EnableWabaInteractions(request *EnableWabaInteractionsRequest) (err error) {
	payload := new(EnableDisableWabaInteractions)
	payload.StatusTemplate = true
	payload.Url = request.Url

	return s.enableOrDisableWabaInteraction(payload)
}

func (s *service) DisableWabaInteractions(request *DisableWabaInteractionsRequest) (err error) {
	payload := new(EnableDisableWabaInteractions)
	payload.StatusTemplate = false
	payload.Url = request.Url

	return s.enableOrDisableWabaInteraction(payload)
}

func (s *service) enableOrDisableWabaInteraction(request *EnableDisableWabaInteractions) (err error) {
	var resp *http.Response

	uri := fmt.Sprint(s.baseUrl, QONTAK_WABA_INTERACTION_URI)
	if resp, err = s.client.Execute(http.MethodPut, s.token, uri, request.ToJSON(), pkg.CONTENT_TYPE_JSON); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to enable waba interactions")
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
