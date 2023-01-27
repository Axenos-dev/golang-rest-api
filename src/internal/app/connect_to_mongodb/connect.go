package connect_to_mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongo_collection struct {
	SuccessfulConnection bool
	Collection           mongo.Collection
}

type ConnectionConfig struct {
	uri             string
	db_name         string
	collection_name string
}

func (config *ConnectionConfig) CreateConnetionConfig(URI string, db_name string, collection_name string) ConnectionConfig {
	config.uri = URI
	config.db_name = db_name
	config.collection_name = collection_name

	return *config
}

func CreateConnection(config ConnectionConfig) mongo_collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.uri))

	if err != nil {
		return mongo_collection{false, mongo.Collection{}}
	}

	client.Connect(context.TODO())

	db := client.Database(config.db_name)

	collection := db.Collection(config.collection_name)

	return mongo_collection{true, *collection}
}

func (collection *mongo_collection) InsertData(data interface{}) {
	collection.Collection.InsertOne(context.TODO(), data)
}

func (collection *mongo_collection) FindData(filter interface{}) *mongo.SingleResult {
	return collection.Collection.FindOne(context.TODO(), filter)
}

func (collection *mongo_collection) FindAllData(filter_value string, operand string, value any) (*mongo.Cursor, error) {
	return collection.Collection.Find(context.TODO(), bson.M{filter_value: bson.M{operand: value}})
}

func (collection *mongo_collection) FindAllDataWithOptions(options options.FindOptions) (*mongo.Cursor, error) {
	return collection.Collection.Find(context.TODO(), bson.M{}, &options)
}

func (collection *mongo_collection) UpdateData(filter interface{}, operand string, value_to_change string, new_value any) error {
	_, err := collection.Collection.UpdateOne(context.TODO(), filter, bson.D{{Key: operand, Value: bson.D{{Key: value_to_change, Value: new_value}}}})

	return err
}

func (collection *mongo_collection) DeleteData(filter interface{}) {
	collection.Collection.DeleteOne(context.TODO(), filter)
}
