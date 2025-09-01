package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func Json(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}
