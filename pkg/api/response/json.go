package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	if body == nil {
		return
	}

	enc, err := json.Marshal(body)
	if err != nil {
		return
	}
	w.Header().Add("content-type", "application/json")
	_, _ = w.Write(enc)
}
