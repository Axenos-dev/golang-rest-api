package recoverpassword

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/lithammer/shortuuid"
)

func Generate_Code() string {
	code := firstN(shortuuid.New(), 6)

	return strings.ToUpper(code)
}

func Generate_Token(code string) string {
	b := make([]byte, 4)
	rand.Read(b)

	return strings.ToLower(code + fmt.Sprintf("%x", b))
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
