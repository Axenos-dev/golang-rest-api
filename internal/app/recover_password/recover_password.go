package recoverpassword

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Recover_Password(token string, new_password string, confirm_password string, collection mongo.Collection) bson.M {
	if len([]rune(confirm_password)) < 6 && len([]rune(new_password)) < 6 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if strings.Contains(new_password, " ") || strings.Contains(confirm_password, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Incorrect password",
		}
	}

	data := bson.M{}

	if err := collection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&data); err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "Incorrect token",
		}
	}

	if data["is_valid"] != true {
		return bson.M{
			"Code":        "204",
			"Description": "Your code is not verified",
		}
	}

	if !Check_Code_Date(data) {
		collection.DeleteOne(context.TODO(), bson.M{"token": token})
		return bson.M{
			"Code":        "204",
			"Description": "Time of code expired",
		}
	}

	if strings.Compare(new_password, confirm_password) != 0 {
		return bson.M{
			"Code":        "204",
			"Description": "Passwords do not match",
		}
	}

	Change_Password(data["email"].(string), new_password)

	collection.DeleteOne(context.TODO(), bson.M{"token": token})

	return bson.M{
		"Code":        "200",
		"Description": "Password changed",
	}
}
