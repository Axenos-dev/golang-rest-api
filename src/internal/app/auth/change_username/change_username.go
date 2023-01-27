package change_username

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

type filter struct {
	Uuid string
}

func ChangeUsername(data models.ChangeUsernameRequest) models.ChangeUsernameResponse {
	var response models.ChangeUsernameResponse

	config := config_auth_db.GetConfig("users")
	collection := connect_to_mongodb.CreateConnection(config)

	if validate_data.ContainsWhiteSpaces(data.New_Username) || validate_data.ContainsWhiteSpaces(data.Uuid) {
		response.CreateResponse(204, "Data must not contain spaces")
		return response
	}

	if validate_data.CheckLength(3, data.New_Username) {
		response.CreateResponse(204, "New username is too short")
		return response
	}

	if err := collection.UpdateData(filter{Uuid: data.Uuid}, "$set", "username", data.New_Username); err != nil {
		response.CreateResponse(204, "Failed to set username")
		return response
	}

	response.CreateResponse(200, "Username changed")
	return response
}
