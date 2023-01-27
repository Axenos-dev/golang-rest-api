package get_games

import (
	"context"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
	"mazano-server/src/internal/app/recommendations/rec_db_config"
	decode_request "mazano-server/src/internal/app/request_decoder"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGames(data models.GetGamesRequest) models.GetGamesResponse {
	var response models.GetGamesResponse

	if data.Limit < 10 {
		response.CreateResponse(204, "Limit must be greater than 10", nil)
		return response
	}

	if data.Bias < 0 {
		response.CreateResponse(204, "Bias must be greater than 0", nil)
		return response
	}

	if data.Limit > 1000 {
		response.CreateResponse(204, "Maximum limit is 1000", nil)
		return response
	}

	if data.Bias+data.Limit > 10000 {
		response.CreateResponse(204, "Incorrect limit", nil)
		return response
	}

	config := rec_db_config.GetConfigAPI("Games-API")
	collection := connect_to_mongodb.CreateConnection(config)

	var film_results []models.Game

	cursor, _ := collection.FindAllDataWithOptions(*options.Find().SetSkip(int64(data.Bias)).SetLimit(int64(data.Limit)))

	cursor.All(context.TODO(), &film_results)

	response.CreateResponse(200, "Results found", decode_request.SnakeCaseEncode(film_results))

	return response
}
