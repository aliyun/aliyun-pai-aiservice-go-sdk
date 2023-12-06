package model

type AigcCreatRequest struct {
	Action        string                 `json:"action"`
	ModelId       string                 `json:"model_id"`
	ModelIds      []string               `json:"model_ids"`
	ModelName     string                 `json:"model_name,omitempty"`
	TemplateImage string                 `json:"template_image,omitempty"`
	Configure     map[string]interface{} `json:"configure,omitempty"`
}
