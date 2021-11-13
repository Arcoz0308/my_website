package api

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	b, _ := json.Marshal(struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}{
		message,
		code,
	})
	w.Write(b)
}
