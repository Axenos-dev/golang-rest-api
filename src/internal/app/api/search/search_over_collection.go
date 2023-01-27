package search

import (
	config_db_api "mazano-server/src/internal/app/api/config_db"
	"mazano-server/src/internal/app/connect_to_mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

func Search_Over_Collection(collection_name, filter_value string, query_string string) (*mongo.Cursor, error) {
	config := config_db_api.GetConfig(collection_name)
	collection := connect_to_mongodb.CreateConnection(config)

	return collection.FindAllData(filter_value, "$regex", query_string)
}
