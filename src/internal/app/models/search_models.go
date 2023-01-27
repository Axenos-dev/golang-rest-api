package models

type Find_Results struct {
	FilmResults   any
	SeriesResults any
	GamesResults  any
}

type Film struct {
	Id                int
	Adult             bool
	Backdrop_path     string
	Genre_ids         []int
	Original_language string
	Original_title    string
	Overview          string
	Popularity        float64
	Poster_path       string
	Release_date      string
	Title             string
	Video             bool
	Vote_average      float64
	Vote_count        int
}

type Series struct {
	Id                int
	Adult             bool
	Backdrop_path     string
	Genre_ids         []int
	Original_language string
	Original_name     string
	Origin_country    []string
	Overview          string
	Popularity        float64
	Poster_path       string
	First_air_date    string
	Name              string
	Vote_average      float64
	Vote_count        int
}

type Game struct {
	Id               int
	Background_image string
	Genres           []string
	Platforms        []string
	Language         string
	Title            string
	Rating           float64
}

type SearchRequest struct {
	Query_String string
}

type SearchResponse struct {
	Code        int16
	Description string
	Results     any
}

func (response *SearchResponse) CreateResponse(code int16, description string, results any) {
	response.Code = code
	response.Description = description
	response.Results = results
}
