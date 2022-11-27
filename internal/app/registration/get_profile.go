package registration

import (
	"context"
	"strings"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get_Profile(uuid string, collection mongo.Collection) bson.M {
	if len([]rune(uuid)) < 1 {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	if strings.Contains(uuid, " ") {
		return bson.M{
			"Code":        "204",
			"Description": "Invalid data",
		}
	}

	var user bson.M

	if err := collection.FindOne(context.TODO(), bson.M{"uuid": uuid}).Decode(&user); err != nil {
		return bson.M{
			"Code":        "204",
			"Description": "No such user with this uuid",
		}
	}

	var favorite_genres []string
	mapstructure.Decode(user["favorite_genres"], &favorite_genres)

	return bson.M{
		"Code":            "200",
		"Username":        user["username"].(string),
		"Email":           user["email"].(string),
		"Avatar":          user["avatar"].(string),
		"Favorite_Genres": favorite_genres,
		"Description":     "User data received",
	}

}
