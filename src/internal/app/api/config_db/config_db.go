package config_db_api

import "mazano-server/src/internal/app/connect_to_mongodb"

func GetConfig(collection_name string) connect_to_mongodb.ConnectionConfig {
	var config connect_to_mongodb.ConnectionConfig
	return config.CreateConnetionConfig("URL to db", "Mazano-API", collection_name)
}
