package models

type LoginUserRequestData struct {
	Email    string
	Password string
}

type LoginUserResponseData struct {
	Code        int16
	Description string
	Uuid        string
}

func (response *LoginUserResponseData) CreateResponse(code int16, description string, uuid string) {
	response.Code = code
	response.Description = description
	response.Uuid = uuid
}
