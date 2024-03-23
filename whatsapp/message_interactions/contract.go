package message_interactions

type Service interface {
	EnableMessageInteractions(request *EnableMessageInteractionsRequest) (err error)
	DisableMessageInteractions(request *DisableMessageInteractionsRequest) (err error)
}
