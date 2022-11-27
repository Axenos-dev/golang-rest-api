package recoverpassword

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Save_Code(code string, email string, token string, collection mongo.Collection) {

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}); err != nil {
		collection.DeleteOne(context.TODO(), bson.M{"email": email})
	}

	data := bson.M{
		"email":    email,
		"code":     code,
		"token":    token,
		"is_valid": false,
		"time": bson.M{
			"Month":  time.Now().Month().String(),
			"Day":    time.Now().Day(),
			"Hour":   time.Now().Hour(),
			"Minute": time.Now().Minute(),
		},
	}

	collection.InsertOne(context.TODO(), data)
}
