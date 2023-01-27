package search

import (
	"context"
	"mazano-server/src/internal/app/models"
	decode_request "mazano-server/src/internal/app/request_decoder"
)

func Search(data models.SearchRequest) models.SearchResponse {
	var response models.SearchResponse

	if Validate_Query(data.Query_String) {
		response.CreateResponse(204, "Invalid query", nil)
		return response
	}

	query_regex := Query_To_Regex(data.Query_String)

	var games_results []models.Game
	var series_results []models.Series
	var films_results []models.Film

	games_cursor, _ := Search_Over_Collection("Games-API", "title", query_regex)
	series_cursor, _ := Search_Over_Collection("Series-API", "name", query_regex)
	films_cursor, _ := Search_Over_Collection("Films-API", "title", query_regex)

	games_cursor.All(context.TODO(), &games_results)
	series_cursor.All(context.TODO(), &series_results)
	films_cursor.All(context.TODO(), &films_results)

	results := models.Find_Results{
		GamesResults:  decode_request.SnakeCaseEncode(games_results),
		FilmResults:   decode_request.SnakeCaseEncode(films_results),
		SeriesResults: decode_request.SnakeCaseEncode(series_results),
	}

	response.CreateResponse(200, "Results found", decode_request.SnakeCaseEncode(results))
	return response
}
