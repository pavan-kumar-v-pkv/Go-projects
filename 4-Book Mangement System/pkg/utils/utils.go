package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the provided interface
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
