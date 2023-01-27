package config_auth_db

import "mazano-server/src/internal/app/connect_to_mongodb"

func GetConfig(collection_name string) connect_to_mongodb.ConnectionConfig {
	var config connect_to_mongodb.ConnectionConfig
	return config.CreateConnetionConfig("URL to db", "mazano", collection_name)
}
