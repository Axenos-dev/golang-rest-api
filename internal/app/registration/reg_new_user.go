package registration

import (
	"context"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register_new_user(username string, password string, email string, collection mongo.Collection) bson.M {

	Uuid := uuid.New().String()

	if len([]rune(username)) < 1 && len([]rune(password)) < 1 && len([]rune(email)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if !ValidMailAddress(email) || strings.Contains(email, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid email address",
		}
	}

	if strings.Contains(password, " ") || len([]rune(password)) < 6 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid password",
		}
	}

	if strings.Contains(username, " ") || len([]rune(username)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid username",
		}
	}

	var user bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user); err == nil {
		return bson.M{
			"Code":        "204",
			"Description": "User with this username already exist",
		}
	}

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err == nil {
		return bson.M{
			"Code":        "204",
			"Description": "User with this email already exist",
		}
	}

	encode_pass, err := HashPassword(password)
	if err != nil {
		return bson.M{
			"Code":        "500",
			"Description": "Internal server error",
		}
	}

	var favorite_genres []string

	data := bson.M{
		"username":        username,
		"password":        encode_pass,
		"email":           email,
		"uuid":            Uuid,
		"avatar":          "",
		"favorite_genres": favorite_genres,
	}

	collection.InsertOne(context.TODO(), data)

	return bson.M{
		"Code":        "200",
		"Uuid":        Uuid,
		"Description": "Successful registration",
	}
}

func ValidMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}
