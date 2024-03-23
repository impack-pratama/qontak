package whatsapp

import (
	"github.com/impack-pratama/qontak/pkg"
	"github.com/impack-pratama/qontak/whatsapp/broadcasts"
	"github.com/impack-pratama/qontak/whatsapp/file_uploader"
	"github.com/impack-pratama/qontak/whatsapp/message_interactions"
	"github.com/impack-pratama/qontak/whatsapp/waba_interactions"
)

const (
	QONTAK_BASE_URL = "https://service-chat.qontak.com/api/open/v1"
)

type Service interface {
	GetBroadCastService() broadcasts.Service
	GetFileUploaderService() file_uploader.Service
	GetWabaInteractionService() waba_interactions.Service
	GetMessageInteractionsService() message_interactions.Service
}

type service struct {
	broadcastsService          broadcasts.Service
	fileUploaderService        file_uploader.Service
	wabaInteractionService     waba_interactions.Service
	messageInteractionsService message_interactions.Service
}

func (s *service) GetMessageInteractionsService() message_interactions.Service {
	return s.messageInteractionsService
}

func (s *service) GetWabaInteractionService() waba_interactions.Service {
	return s.wabaInteractionService
}

func (s *service) GetFileUploaderService() file_uploader.Service {
	return s.fileUploaderService
}

func (s *service) GetBroadCastService() broadcasts.Service {
	return s.broadcastsService
}

func NewService(token string) Service {
	client := pkg.NewHttp()

	a := new(service)
	a.broadcastsService = broadcasts.NewService(client, token, QONTAK_BASE_URL)
	a.fileUploaderService = file_uploader.NewService(client, token, QONTAK_BASE_URL)
	a.wabaInteractionService = waba_interactions.NewService(client, token, QONTAK_BASE_URL)
	a.messageInteractionsService = message_interactions.NewService(client, token, QONTAK_BASE_URL)
	return a
}
