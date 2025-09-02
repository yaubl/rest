package response

import (
	// ignore this visual error.
	"encoding/json/v2"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func Json(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.MarshalWrite(w, v, nil)
	}
}
