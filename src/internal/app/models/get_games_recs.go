package models

type GetGamesRequest struct {
	Uuid  string
	Bias  int
	Limit int
}

type GetGamesResponse struct {
	Code        int16
	Description string
	Results     any
}

func (response *GetGamesResponse) CreateResponse(code int16, description string, results any) {
	response.Code = code
	response.Description = description
	response.Results = results
}
