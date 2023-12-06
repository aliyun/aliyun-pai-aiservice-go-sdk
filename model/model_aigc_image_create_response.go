package model

type AIGCImageCreateData struct {
	Image    string  `json:"image"`
	CostTime float64 `json:"cost_time"`
}
type AIGCImageCreateResponse struct {
	RequestId string              `json:"request_id,omitempty"`
	Code      string              `json:"code,omitempty"`
	Message   string              `json:"message,omitempty"`
	Data      AIGCImageCreateData `json:"data,omitempty"`
}

func NewAIGCImageCreateResponse() *AIGCImageCreateResponse {
	return &AIGCImageCreateResponse{}
}
