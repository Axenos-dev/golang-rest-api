package confirm_code

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/recover_password/check_code_date"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
	"strings"
)

type find_by_token struct {
	token string
}

func ConfirmCode(data models.ConfirmCodeRequest) models.ConfirmCodeResponse {
	var response models.ConfirmCodeResponse
	config := config_auth_db.GetConfig("recover-password")

	collection := connect_to_mongodb.CreateConnection(config)

	var find_results models.ValidationCode

	if validate_data.CheckLength(6, data.Code) || validate_data.ContainsWhiteSpaces(data.Code) {
		response.CreateResponse(204, "Invalid data", "")
		return response
	}

	if err := collection.FindData(find_by_token{token: data.Token}).Decode(&find_results); err != nil {
		response.CreateResponse(204, "Token not found", "")
		return response
	}

	if !check_code_date.Check_Date(find_results) {
		collection.DeleteData(find_by_token{token: data.Token})

		response.CreateResponse(204, "Time of code expired", "")
		return response
	}

	var code string = strings.ToUpper(data.Code)

	if strings.Compare(code, find_results.Code) != 0 {
		response.CreateResponse(204, "Code is wrong", "")
		return response
	}

	if find_results.Is_Valid {
		response.CreateResponse(204, "Your code is already verified", "")
		return response
	}

	err := collection.UpdateData(find_by_token{token: data.Token}, "$set", "is_valid", true)

	if err != nil {
		response.CreateResponse(500, err.Error(), "")
		return response
	}

	response.CreateResponse(200, "Code valid", data.Token)

	return response
}
