package get_profile

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

func GetProfile(data models.GetProfileRequest) models.GetProfileResponse {
	var response models.GetProfileResponse
	config := config_auth_db.GetConfig("users")

	collection := connect_to_mongodb.CreateConnection(config)

	if validate_data.ContainsWhiteSpaces(data.Uuid) {
		response.CreateResponse(204, "Invalid data", nil)
		return response
	}

	var find_results models.Profile
	if err := collection.FindData(data).Decode(&find_results); err != nil {
		response.CreateResponse(204, "User not found", nil)
		return response
	}

	response.CreateResponse(200, "User found", find_results)
	return response
}
