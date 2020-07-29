package response

import (
	"encoding/json"
	"net/http"
)

// JSON will send a response with a json body
func JSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body == nil {
		return
	}

	enc, err := json.Marshal(body)
	if err != nil {
		return
	}
	_, _ = w.Write(enc)
}
