package decode_request

import (
	"reflect"
	"strings"
	"unicode/utf8"
)

func SnakeCaseEncode(i interface{}) any {
	rt, rv := reflect.TypeOf(i), reflect.ValueOf(i)

	if rt.Kind() == reflect.Slice {
		ret := make([]interface{}, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			ret[i] = rv.Index(i).Interface()
		}

		var result []interface{}

		for _, elem := range ret {
			result = append(result, SnakeCaseEncode(elem))
		}

		return result
	}

	if rt.Kind() == reflect.Ptr {
		i := reflect.Indirect(rv).Interface()
		rt, rv = reflect.TypeOf(i), reflect.ValueOf(i)
	}

	out := make(map[string]interface{}, rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		if rt.Field(i).Tag.Get("json") == "-" {
			continue
		}

		if rt.Field(i).Tag.Get("json") == "_" {
			continue
		}

		if strings.Contains(rt.Field(i).Tag.Get("json"), "omitempty") && rv.Field(i).IsZero() {
			continue
		}

		k := snakeCase(rt.Field(i).Name)

		out[k] = rv.Field(i).Interface()
	}

	return out
}

func snakeCase(s string) string {
	out := make([]rune, 0, utf8.RuneCountInString(s))

	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 32

			if i > 0 {
				out = append(out, '_')
			}
		}

		out = append(out, r)
	}

	return string(out)
}
