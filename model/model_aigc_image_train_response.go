package model

type AIGCImageTrainData struct {
	JobId   int32  `json:"job_id"`
	ModelId string `json:"model_id"`
}
type AIGCImageTrainResponse struct {
	RequestId string             `json:"request_id,omitempty"`
	Code      string             `json:"code,omitempty"`
	Message   string             `json:"message,omitempty"`
	Data      AIGCImageTrainData `json:"data,omitempty"`
}

func NewAIGCImageTrainResponse() *AIGCImageTrainResponse {
	return &AIGCImageTrainResponse{}
}
