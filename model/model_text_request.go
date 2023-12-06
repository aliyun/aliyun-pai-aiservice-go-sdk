package model

type TextRequest struct {
	Action    string                 `json:"action"`
	Text      string                 `json:"text"`
	ModelName string                 `json:"model_name,omitempty"`
	Configure map[string]interface{} `json:"configure,omitempty"`
}
