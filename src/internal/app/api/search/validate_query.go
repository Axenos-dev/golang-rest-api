package search

func Validate_Query(query_string string) bool {
	spaces := 0

	for _, char := range query_string {
		if string(char) == " " {
			spaces++
		}
	}

	return len(query_string)-spaces < 1
}
