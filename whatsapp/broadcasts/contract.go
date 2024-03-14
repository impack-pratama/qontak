package broadcasts

import "github.com/impack-pratama/qontak/whatsapp/broadcasts/direct_message"

type Service interface {
	// SendDirectMessage sends a direct message to the WhatsApp API
	SendDirectMessage(request *direct_message.SendDirectMessageRequest) (response *direct_message.SendDirectMessageResponse, err error)
}
