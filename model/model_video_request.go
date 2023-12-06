package model

type VideoRequest struct {
	Action     string                 `json:"action"`
	VideoUrl   string                 `json:"video_url"`
	VideoTitle string                 `json:"video_title,omitempty"`
	ModelName  string                 `json:"model_name,omitempty"`
	Text       string                 `json:"text,omitempty"`
	Images     []string               `json:"images,omitempty"`
	Configure  map[string]interface{} `json:"configure,omitempty"`
}
