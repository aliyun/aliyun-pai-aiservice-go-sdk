package model

const (
	RESPONSE_OK = "OK"
)

type Response struct {
	RequestId string                 `json:"request_id,omitempty"`
	Code      string                 `json:"code,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

func NewResponse() *Response {
	return &Response{
		Data: make(map[string]interface{}, 0),
	}
}
