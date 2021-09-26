package api

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	b, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		message,
	})
	w.Write(b)
}
