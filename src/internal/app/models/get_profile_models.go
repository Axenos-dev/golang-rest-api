package models

type Profile struct {
	Username        string
	Email           string
	Avatar          string
	Favorite_Genres []string
}

type GetProfileRequest struct {
	Uuid string
}

type GetProfileResponse struct {
	Code        int16
	Description string
	Profile     any
}

func (response *GetProfileResponse) CreateResponse(code int16, description string, profile any) {
	response.Code = code
	response.Description = description
	response.Profile = profile
}
