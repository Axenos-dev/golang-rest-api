package recoverpassword

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
)

type Date struct {
	Month  string
	Day    int
	Hour   int
	Minute int
}

func Check_Code_Date(data bson.M) bool {
	var date Date
	mapstructure.Decode(data["time"], &date)

	time_now := time.Now()

	if time_now.Month().String() != date.Month || time_now.Day() != date.Day || time_now.Hour() != date.Hour || time_now.Minute()-date.Minute > 5 {
		return false
	}
	return true
}
