package check_code_date

import (
	"mazano-server/src/internal/app/models"
	"time"
)

func Check_Date(data models.ValidationCode) bool {
	time_now := time.Now()
	code_date := data.Time

	if code_date.Month != time_now.Month().String() || code_date.Day != time_now.Day() || time_now.Hour()-code_date.Hour > 1 || time_now.Minute()-code_date.Minute > 5 {
		return false
	}

	return true
}
