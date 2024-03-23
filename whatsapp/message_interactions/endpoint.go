package message_interactions

type EnableMessageInteractionsRequest struct {
	ReceiveMessageFromCustomer bool
	ReceiveMessageFromAgent    bool
	Url                        string
}

type DisableMessageInteractionsRequest struct {
	ReceiveMessageFromCustomer bool
	ReceiveMessageFromAgent    bool
	Url                        string
}
