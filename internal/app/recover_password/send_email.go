package recoverpassword

import (
	"mazano-server/internal/app/registration"
	"net/smtp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendEmail(email string, code string, token string, collection mongo.Collection) bson.M {
	if len([]rune(email)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if !registration.ValidMailAddress(email) || strings.Contains(email, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid email",
		}
	}

	if !Does_User_Exist(email, "users") {
		return bson.M{
			"Code":        "204",
			"Description": "No such user in database",
		}
	}

	// Sender data.
	from := "mazanoserver@gmail.com"
	password := "ovobypmvgqcquygm"

	// Receiver email address.
	to := []string{
		email,
	}

	// Message.
	message := []byte(code)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	// Sending email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		return bson.M{
			"Code":        "204",
			"Description": err.Error(),
		}
	}

	Save_Code(code, email, token, collection)

	return bson.M{
		"Code":        "200",
		"Token":       token,
		"Description": "Email sent",
	}
}
