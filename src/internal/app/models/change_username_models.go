package models

type ChangeUsernameRequest struct {
	Uuid         string
	New_Username string
}

type ChangeUsernameResponse struct {
	Code        int16
	Description string
}

func (response *ChangeUsernameResponse) CreateResponse(code int16, description string) {
	response.Code = code
	response.Description = description
}
