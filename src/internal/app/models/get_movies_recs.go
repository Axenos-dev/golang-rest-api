package models

type GetMoviesRequest struct {
	Uuid  string
	Bias  int
	Limit int
}

type GetMoviesResponse struct {
	Code        int16
	Description string
	Results     any
}

func (response *GetMoviesResponse) CreateResponse(code int16, description string, results any) {
	response.Code = code
	response.Description = description
	response.Results = results
}
