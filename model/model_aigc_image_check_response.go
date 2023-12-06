package model

type AIGCImageCheckResult struct {
	Code    int    `json:"code"`
	Frontal bool   `json:"frontal"`
	Message string `json:"message"`
	Url     string `json:"url"`
}
type AIGCImageCheckData struct {
	RequestId    string                 `json:"request_id"`
	Images       []string               `json:"images"`
	CostTime     float64                `json:"cost_time"`
	CheckResults []AIGCImageCheckResult `json:"check_results"`
}
type AIGCImageCheckResponse struct {
	RequestId string             `json:"request_id,omitempty"`
	Code      string             `json:"code,omitempty"`
	Message   string             `json:"message,omitempty"`
	Data      AIGCImageCheckData `json:"data,omitempty"`
}

func NewAIGCImageCheckResponse() *AIGCImageCheckResponse {
	return &AIGCImageCheckResponse{}
}
