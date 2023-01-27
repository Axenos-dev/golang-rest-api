package send_email

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/auth/recover_password/save_code"
	"mazano-server/src/internal/app/auth/recover_password/sender_config"
	"mazano-server/src/internal/app/auth/validate_data"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

type find_data struct {
	Email string
}

func Send_Email(data models.SendEmailRequest) models.SendEmailResponse {
	var response models.SendEmailResponse
	var config sender_config.SMTPConfig
	var validationCode models.ValidationCode

	config_db := config_auth_db.GetConfig("users")
	collection := connect_to_mongodb.CreateConnection(config_db)

	if !validate_data.ValidateEmail(data.Email) {
		response.CreateResponse(204, "Invalid email", "")

		return response
	}

	var find_results models.User
	if err := collection.FindData(find_data{Email: data.Email}).Decode(&find_results); err != nil {
		response.CreateResponse(204, "No such user with this email", "")

		return response
	}

	validationCode.GenerateCode(data.Email)

	// Sender data.
	config.CreateConfig("mazanoserver@gmail.com", "ovobypmvgqcquygm")

	// Receiver email address.
	to := []string{
		data.Email,
	}

	// Message.
	message := []byte(validationCode.Code)

	// Authentication.
	auth := sender_config.CreateConnection(config)

	err := sender_config.Send_Mail(auth, config, to, message)

	if err != nil {
		response.CreateResponse(204, err.Error(), "")
	}

	save_code.Save_Code(validationCode)

	response.CreateResponse(200, "Email sent", validationCode.Token)
	return response
}
