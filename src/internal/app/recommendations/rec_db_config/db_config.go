package rec_db_config

import "mazano-server/src/internal/app/connect_to_mongodb"

func GetConfigAPI(collection_name string) connect_to_mongodb.ConnectionConfig {
	var config connect_to_mongodb.ConnectionConfig
	return config.CreateConnetionConfig("URL to db", "Mazano-API", collection_name)
}
