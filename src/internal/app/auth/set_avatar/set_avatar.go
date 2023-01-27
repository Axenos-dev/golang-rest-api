package set_avatar

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

type filter struct {
	Uuid string
}

func SetAvatar(data models.SetAvatarRequest) models.SetAvatarResponse {
	var response models.SetAvatarResponse

	config := config_auth_db.GetConfig("users")
	collection := connect_to_mongodb.CreateConnection(config)

	if validate_data.ContainsWhiteSpaces(data.New_Avatar) || validate_data.ContainsWhiteSpaces(data.Uuid) {
		response.CreateResponse(204, "Data must not contain spaces")
		return response
	}

	if err := collection.UpdateData(filter{Uuid: data.Uuid}, "$set", "avatar", data.New_Avatar); err != nil {
		response.CreateResponse(204, "Failed to set avatar")
		return response
	}

	response.CreateResponse(200, "Avatar changed")
	return response
}
