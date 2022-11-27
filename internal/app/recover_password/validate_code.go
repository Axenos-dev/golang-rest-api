package recoverpassword

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate_Code(token string, code string, collection mongo.Collection) bson.M {
	if len([]rune(token)) < 1 || len([]rune(code)) < 1 {
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "Invalid data",
		}
	}

	if strings.Contains(token, " ") || strings.Contains(code, " ") {
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "Data must not contain spaces",
		}
	}

	data := bson.M{}

	if err := collection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&data); err != nil {
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "This token does not exist",
		}
	}

	if data["code"] != code {
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "Incorrect validation code",
		}
	}

	if !Check_Code_Date(data) {
		collection.DeleteOne(context.TODO(), bson.M{"token": token})
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "Time of code expired",
		}
	}

	if data["is_valid"] == true {
		return bson.M{
			"Code":        "204",
			"Valid":       true,
			"Token":       token,
			"Description": "Your code is already verified",
		}
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"token": token}, bson.D{{"$set", bson.D{{"is_valid", true}}}})
	if err != nil {
		return bson.M{
			"Code":        "204",
			"Valid":       false,
			"Description": "Unexpected error",
		}
	}

	return bson.M{
		"Code":        "200",
		"Token":       token,
		"Valid":       true,
		"Description": "Code is valid",
	}
}
