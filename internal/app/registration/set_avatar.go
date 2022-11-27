package registration

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Set_Avatar(avatar_url string, uuid string, collection mongo.Collection) bson.M {

	if len([]rune(avatar_url)) < 3 || len([]rune(uuid)) < 3 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if strings.Contains(avatar_url, " ") || strings.Contains(uuid, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Data must not contain spaces",
		}
	}

	var user bson.M

	if err := collection.FindOne(context.TODO(), bson.M{"uuid": uuid}).Decode(&user); err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "Such user does not exist",
		}
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"uuid": uuid}, bson.D{{"$set", bson.D{{"avatar", avatar_url}}}})
	if err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "Unexpected error",
		}
	}

	return bson.M{
		"Code":        "200",
		"Description": "Avatar changed",
	}
}
