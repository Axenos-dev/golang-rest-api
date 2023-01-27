package models

type SetAvatarRequest struct {
	Uuid       string
	New_Avatar string
}

type SetAvatarResponse struct {
	Code        int16
	Description string
}

func (response *SetAvatarResponse) CreateResponse(code int16, description string) {
	response.Code = code
	response.Description = description
}
