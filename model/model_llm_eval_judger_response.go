package model

type LLMEvalJudgerResponse struct {
	RequestId string                 `json:"request_id,omitempty"`
	Code      string                 `json:"code,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Data      ChatCompletionResponse `json:"data,omitempty"`
}

type ChatCompletionResponse struct {
	Id      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Index        int    `json:"index"`
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

func NewLLMEvalJudgerResponse() *LLMEvalJudgerResponse {
	return &LLMEvalJudgerResponse{}
}
