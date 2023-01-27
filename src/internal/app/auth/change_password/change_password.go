package change_password

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	hashpassword "mazano-server/src/internal/app/auth/hash_password"
	"mazano-server/src/internal/app/auth/recover_password/check_code_date"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

type filter_recover_password struct {
	token string
}

type filter_users struct {
	email string
}

func ChangePassword(data models.ChangePasswordRequest) models.ChangePasswordResponse {
	var response models.ChangePasswordResponse

	if validate_data.CheckLength(6, data.New_Password) || validate_data.CheckLength(6, data.Confirm_New_Password) {
		response.CreateResponse(204, "Password must be minimum 6 in length")
		return response
	}

	if validate_data.ContainsWhiteSpaces(data.New_Password) || validate_data.ContainsWhiteSpaces(data.Confirm_New_Password) {
		response.CreateResponse(204, "Password must not contain spaces")
		return response
	}

	config_recover_password := config_auth_db.GetConfig("recover-password")
	collection_recover_password := connect_to_mongodb.CreateConnection(config_recover_password)

	var find_results models.ValidationCode
	if err := collection_recover_password.FindData(filter_recover_password{token: data.Token}).Decode(&find_results); err != nil {
		response.CreateResponse(204, "Token not found")
		return response
	}

	if !find_results.Is_Valid {
		response.CreateResponse(204, "You didn`t verify code")
		return response
	}

	if !check_code_date.Check_Date(find_results) {
		collection_recover_password.DeleteData(filter_recover_password{token: data.Token})

		response.CreateResponse(204, "Time of code expired")
		return response
	}

	if data.New_Password != data.Confirm_New_Password {
		response.CreateResponse(204, "Passwords doesn`t match")
		return response
	}

	config_users := config_auth_db.GetConfig("users")
	collection_users := connect_to_mongodb.CreateConnection(config_users)

	pass_hash, _ := hashpassword.HashPassword(data.New_Password)

	err := collection_users.UpdateData(filter_users{email: find_results.Email}, "$set", "password", pass_hash)

	if err != nil {
		response.CreateResponse(500, err.Error())
		return response
	}

	collection_users.DeleteData(filter_recover_password{token: data.Token})

	response.CreateResponse(200, "Password changed")
	return response
}
