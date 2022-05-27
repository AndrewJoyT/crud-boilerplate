package email

type TemplateEmailRequest struct {
	TemplateName string      `json:"template_name"`
	Recipients   []string    `json:"recipients"`
	Subject      string      `json:"subject"`
	Data         interface{} `json:"data"`
}

type RegisterConfirmationRequest struct {
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Url      string `json:"url"`
}
