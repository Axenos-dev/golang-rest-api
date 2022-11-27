package recoverpassword

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoRequestRecoverPassword(method string, data bson.M) bson.M {
	URI := "Mongo URI"
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		return bson.M{
			"Code":        "500",
			"Description": "Internal server error",
		}
	}

	client.Connect(context.TODO())

	db := client.Database("mazano")
	collection := db.Collection("recover-password")

	switch method {
	case "Send-Email":
		{
			code := Generate_Code()
			token := Generate_Token(code)

			return SendEmail(data["email"].(string), code, token, *collection)
		}

	case "Code-Validation":
		{
			return Validate_Code(data["token"].(string), data["code"].(string), *collection)
		}

	case "Recover-Password":
		{
			return Recover_Password(data["token"].(string), data["new_password"].(string), data["confirm_password"].(string), *collection)
		}
	}

	defer client.Disconnect(context.TODO())
	return bson.M{
		"000": "Allo",
	}
}
