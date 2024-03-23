package waba_interactions

type Service interface {
	EnableWabaInteractions(request *EnableWabaInteractionsRequest) (err error)
	DisableWabaInteractions(request *DisableWabaInteractionsRequest) (err error)
}
