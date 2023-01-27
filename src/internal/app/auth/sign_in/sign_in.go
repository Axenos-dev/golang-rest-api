package sign_in

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

func Sign_In(data models.LoginUserRequestData) models.LoginUserResponseData {
	var response models.LoginUserResponseData
	var find_user_query models.SearchForUserByEmail

	if validate_data.ContainsWhiteSpaces(data.Email) || validate_data.ContainsWhiteSpaces(data.Password) {
		response.CreateResponse(204, "Data must not contain spaces", "")

		return response
	}

	if !validate_data.ValidateEmail(data.Email) {
		response.CreateResponse(204, "Invalid email", "")

		return response
	}

	config := config_auth_db.GetConfig("users")
	collection := connect_to_mongodb.CreateConnection(config)

	find_user_query.CreateQuery(data.Email)
	var find_results models.User

	if err := collection.FindData(find_user_query).Decode(&find_results); err != nil {
		response.CreateResponse(204, "User with this email doesnt exist", "")

		return response
	}

	if !find_results.ComparePasswords(data.Password) {
		response.CreateResponse(204, "Wrong password", "")

		return response
	}

	response.CreateResponse(200, "Successful login", find_results.Uuid)

	return response
}
