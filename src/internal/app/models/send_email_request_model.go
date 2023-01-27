package models

type SendEmailRequest struct {
	Email string
}

type SendEmailResponse struct {
	Code        int16
	Description string
	Token       string
}

func (response *SendEmailResponse) CreateResponse(code int16, description string, token string) {
	response.Code = code
	response.Description = description
	response.Token = token
}
