package account

import "github.com/gorilla/mux"

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/verify-email/{key}", verifyEmail).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET", "POST")
}
