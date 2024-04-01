package model

type LlmEvalJudgerRequest struct {
	Action    string                         `json:"action"`
	ModelName string                         `json:"model_name,omitempty"`
	Model     string                         `json:"model"`
	Messages  []LlmEvalJudgerRequestMessages `json:"messages"`
}
