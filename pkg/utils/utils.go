package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// since helper function use blank interface so we can pass in any interface
// we create for this code we will be passing book interface
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
