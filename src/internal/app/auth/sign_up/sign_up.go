package sign_up

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

func Sign_Up(data models.NewUserRequestData) models.NewUserResponse {
	var response models.NewUserResponse
	var user models.NewUser
	var overlap models.SearchForUserByEmail

	if validate_data.ContainsWhiteSpaces(data.Username) || validate_data.ContainsWhiteSpaces(data.Password) || validate_data.ContainsWhiteSpaces(data.Email) {
		response.CreateResponse(204, "Data must not contain spaces", user)

		return response
	}

	if validate_data.CheckLength(4, data.Username) || validate_data.CheckLength(4, data.Password) {
		response.CreateResponse(204, "Password and username must have minimum 4 symbols in it", user)

		return response
	}

	if !validate_data.ValidateEmail(data.Email) {
		response.CreateResponse(204, "Invalid email", user)

		return response
	}

	config := config_auth_db.GetConfig("users")

	collection := connect_to_mongodb.CreateConnection(config)

	if !collection.SuccessfulConnection {
		response.CreateResponse(502, "Error with connecting to db", user)

		return response
	}

	user.CreateUser(data)
	overlap.CreateQuery(user.Email)

	var find_results interface{}
	if err := collection.FindData(overlap).Decode(&find_results); err == nil {
		response.CreateResponse(204, "User with this email already exists", user)

		return response
	}

	collection.InsertData(user)

	response.CreateResponse(200, "Successful registration", user)

	return response
}
