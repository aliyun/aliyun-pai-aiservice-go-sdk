package model

type LlmEvalJudgerRequestMessages struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
