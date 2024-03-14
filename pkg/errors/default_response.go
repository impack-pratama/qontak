package errors

type Error struct {
	Code     int      `json:"code,omitempty"`
	Messages []string `json:"messages,omitempty"`
}
type DefaultErrorResponse struct {
	Status string `json:"status,omitempty"`
	Error  Error  `json:"error,omitempty"`
}
