package model

type AsyncJobResult struct {
	State   int    `json:"state"`
	AppId   string `json:"app_id"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type AsyncJobData struct {
	Job AsyncJobResult `json:"job"`
}
type AsyncJobResponse struct {
	RequestId string       `json:"request_id,omitempty"`
	Code      string       `json:"code,omitempty"`
	Message   string       `json:"message,omitempty"`
	Data      AsyncJobData `json:"data,omitempty"`
}

func NewAsyncJobResponse() *AsyncJobResponse {
	return &AsyncJobResponse{}
}
