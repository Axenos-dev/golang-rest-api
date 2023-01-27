package decode_request

import (
	"encoding/json"
	"net/http"
)

func Decode_Request(req *http.Request, data interface{}) {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	dec.Decode(data)
}
