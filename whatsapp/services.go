package whatsapp

import (
	"github.com/impack-pratama/qontak/pkg"
	"github.com/impack-pratama/qontak/whatsapp/broadcasts"
)

const (
	QONTAK_BASE_URL = "https://service-chat.qontak.com/api/open/v1"
)

type Service interface {
	GetBroadCastService() broadcasts.Service
}

type service struct {
	broadcastsService broadcasts.Service
}

func (s *service) GetBroadCastService() broadcasts.Service {
	return s.broadcastsService
}

func NewService(token string) Service {
	client := pkg.NewHttp()

	a := new(service)
	a.broadcastsService = broadcasts.NewService(client, token, QONTAK_BASE_URL)
	return a
}
