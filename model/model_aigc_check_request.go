package model

type AigcCheckRequest struct {
	Action    string                 `json:"action"`
	Images    []string               `json:"images"`
	ModelName string                 `json:"model_name,omitempty"`
	Configure map[string]interface{} `json:"configure,omitempty"`
}
