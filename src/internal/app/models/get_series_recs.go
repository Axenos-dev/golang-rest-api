package models

type GetSeriesRequest struct {
	Uuid  string
	Bias  int
	Limit int
}

type GetSeriesResponse struct {
	Code        int16
	Description string
	Results     any
}

func (response *GetSeriesResponse) CreateResponse(code int16, description string, results any) {
	response.Code = code
	response.Description = description
	response.Results = results
}
