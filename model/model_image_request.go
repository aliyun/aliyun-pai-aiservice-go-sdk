package model

type ImageRequest struct {
	Action    string                 `json:"action"`
	Image     string                 `json:"image"`
	ModelName string                 `json:"model_name,omitempty"`
	Configure map[string]interface{} `json:"configure,omitempty"`
}
