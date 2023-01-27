package models

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type time_now struct {
	Month  string
	Day    int
	Hour   int
	Minute int
}

type ValidationCode struct {
	Code     string
	Email    string
	Token    string
	Is_Valid bool
	Time     time_now
}

func (code *ValidationCode) GenerateCode(email string) {
	now := time.Now()
	timenow := time_now{
		Month:  now.Month().String(),
		Day:    now.Day(),
		Hour:   now.Hour(),
		Minute: now.Minute(),
	}

	code.Email = email
	code.Code = strings.ToUpper(firstN(uuid.New().String(), 6))
	code.Token = generate_token(code.Code)
	code.Time = timenow
	code.Is_Valid = false
}

func generate_token(code string) string {
	bytes := make([]byte, 4)
	rand.Read(bytes)

	return strings.ToLower(code + fmt.Sprintf("%x", bytes))
}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
