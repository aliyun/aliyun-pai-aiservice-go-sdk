package model

type AiCreationRequest struct {
	Action    string                 `json:"action"`
	Content   string                 `json:"content"`
	ModelName string                 `json:"model_name,omitempty"`
	Configure map[string]interface{} `json:"configure,omitempty"`
}
