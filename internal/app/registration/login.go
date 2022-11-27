package registration

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(email string, password string, collection mongo.Collection) bson.M {

	if len([]rune(password)) < 1 && len([]rune(email)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if !ValidMailAddress(email) || strings.Contains(email, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Wrong email address",
		}
	}

	if strings.Contains(password, " ") || len([]rune(password)) < 6 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid password",
		}
	}

	var userLogin bson.M

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&userLogin); err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "User with this email does not exist",
		}
	}
	p := userLogin["password"].(string)

	if !CheckPasswordHash(password, p) {
		return bson.M{
			"Code":        "204",
			"Description": "Incorrect password",
		}
	}

	uuid := userLogin["uuid"].(string)

	return bson.M{
		"Code":        "200",
		"Uuid":        uuid,
		"Description": "Successful login",
	}
}
