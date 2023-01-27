package models

type ChangePasswordRequest struct {
	New_Password         string
	Confirm_New_Password string
	Token                string
}

type ChangePasswordResponse struct {
	Code        int16
	Description string
}

func (response *ChangePasswordResponse) CreateResponse(code int16, description string) {
	response.Code = code
	response.Description = description
}
