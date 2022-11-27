package recoverpassword

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Test struct {
}

func Does_User_Exist(email string, collection_name string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Axenos:68bRJbAoKMOYwkd9@mazano.ovwc4va.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		fmt.Println("1")
		return false
	}

	client.Connect(context.TODO())

	db := client.Database("mazano")
	collection := db.Collection(collection_name)

	var user bson.M

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err == nil {
		return true
	}

	defer client.Disconnect(context.TODO())

	return false
}
