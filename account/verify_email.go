package account

import (
	"github.com/gorilla/mux"
	"net/http"
)

func verifyEmail(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)["key"]
	if value == "" {

	}
}
