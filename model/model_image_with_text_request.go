package model

type ImageWithTextRequest struct {
	Action    string                 `json:"action"`
	Images    []string               `json:"images"`
	Text      string                 `json:"text,omitempty"`
	ModelName string                 `json:"model_name,omitempty"`
	Configure map[string]interface{} `json:"configure,omitempty"`
}
