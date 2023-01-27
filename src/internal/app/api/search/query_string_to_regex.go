package search

import "strings"

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
