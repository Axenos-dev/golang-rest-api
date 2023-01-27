package validate_data

import (
	"net/mail"
	"regexp"
)

func ContainsWhiteSpaces(s string) bool {
	return regexp.MustCompile(`\s`).MatchString(s)
}

func ValidateEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func CheckLength(l int, s string) bool {
	return len(s) < l
}
