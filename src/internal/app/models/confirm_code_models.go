package models

type ConfirmCodeRequest struct {
	Token string
	Code  string
}

type ConfirmCodeResponse struct {
	Code        int16
	Description string
	Token       string
}

func (response *ConfirmCodeResponse) CreateResponse(code int16, description string, token string) {
	response.Code = code
	response.Description = description
	response.Token = token
}
