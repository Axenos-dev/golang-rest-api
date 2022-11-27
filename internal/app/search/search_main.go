package search_main

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Search_MazanoAPI(query_data bson.M) bson.M {
	if Check_For_Spaces(query_data["query_string"].(string)) {
		return bson.M{
			"Code":        "204",
			"Description": "Bad query",
		}
	}

	URI := "Mongo URI"
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		return bson.M{
			"Code":        "404",
			"Description": "Error with connecting to api",
		}
	}

	client.Connect(context.TODO())

	db := client.Database("Mazano-API")

	query_regex := Query_To_Regex(query_data["query_string"].(string))

	films_collection := db.Collection("Films-API")
	series_collection := db.Collection("Series-API")
	games_collection := db.Collection("Games-API")

	var games_results []bson.M
	var films_results []bson.M
	var series_results []bson.M

	games_results = Search_Over_Collection(games_collection, query_regex)
	films_results = Search_Over_Collection(films_collection, query_regex)
	series_results = Search_Over_Collection(series_collection, query_regex)

	return bson.M{
		"Code":        "200",
		"Description": "All possible results found",
		"Results": bson.M{
			"games_results":  games_results,
			"films_results":  films_results,
			"series_results": series_results,
		},
	}
}

func Search_Over_Collection(collection *mongo.Collection, query_string_regex string) []bson.M {
	cursor, err := collection.Find(context.TODO(), bson.M{"title": bson.M{"$regex": query_string_regex}})

	var results []bson.M

	if err != nil {
		fmt.Println(err.Error())

		return []bson.M{}
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())

		return []bson.M{}
	}

	return results
}

func Query_To_Regex(query_string string) string {
	regex := ""

	for _, char := range query_string {
		if string(char) != " " {
			regex += "[" + strings.ToUpper(string(char)) + strings.ToLower(string(char)) + "]"
		} else {
			regex += " "
		}
	}

	return regex
}

func Check_For_Spaces(query_string string) bool {
	spaces := 0

	for _, char := range query_string {
		if string(char) == " " {
			spaces++
		}
	}

	if len(query_string)-spaces < 3 {
		return true
	}

	return false
}
