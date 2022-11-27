package registration

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Change_Username(uuid string, new_username string, collection mongo.Collection) bson.M {

	if len([]rune(uuid)) < 1 || len([]rune(new_username)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if strings.Contains(uuid, " ") || strings.Contains(new_username, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	var user1 bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"username": new_username}).Decode(&user1); err == nil {
		return bson.M{
			"Code":        "204",
			"Description": "User with this username already exist",
		}
	}

	var user2 bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"uuid": uuid}).Decode(&user2); err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "Such user does not exist",
		}
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"uuid": uuid}, bson.D{{"$set", bson.D{{"username", new_username}}}})
	if err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "Unexpected error",
		}
	}

	return bson.M{
		"Code":        "200",
		"Description": "Username changed",
	}
}
