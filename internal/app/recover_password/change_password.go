package recoverpassword

import (
	"context"
	"mazano-server/internal/app/registration"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Change_Password(email string, new_password string) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Axenos:68bRJbAoKMOYwkd9@mazano.ovwc4va.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		return
	}

	client.Connect(context.TODO())

	db := client.Database("mazano")

	collection := db.Collection("users")

	hash_pass, _ := registration.HashPassword(new_password)

	_, err = collection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.D{{"$set", bson.D{{"password", hash_pass}}}})

	if err != nil {
		return
	}
}
