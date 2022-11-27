package registration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoRequest(data bson.M) bson.M {
	URI := "Mongo URI"
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		return bson.M{
			"Code":        "404",
			"Description": "Error with connecting to users database",
		}
	}

	client.Connect(context.TODO())

	db := client.Database("mazano")

	collection := db.Collection("users")

	switch data["request_type"].(string) {
	case "INSERT":
		{
			switch data["method"].(string) {
			case "register":
				{
					return Register_new_user(data["username"].(string), data["password"].(string), data["email"].(string), *collection)
				}
			}
		}

	case "GET":
		{
			switch data["method"].(string) {
			case "login":
				{
					return LoginUser(data["email"].(string), data["password"].(string), *collection)
				}

			case "profile":
				{
					return Get_Profile(data["uuid"].(string), *collection)
				}
			}
		}

	case "SET":
		{
			switch data["method"].(string) {
			case "change-username":
				{
					return Change_Username(data["uuid"].(string), data["new_username"].(string), *collection)
				}

			case "set-avatar":
				{
					return Set_Avatar(data["new_avatar"].(string), data["uuid"].(string), *collection)
				}
			}
		}
	}
	defer client.Disconnect(context.TODO())

	return bson.M{
		"Allo": "000",
	}
}
