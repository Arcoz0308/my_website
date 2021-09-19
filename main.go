package main

import (
	"github.com/arcoz0308/arcoz0308.tech/api"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	utils.LoadConfig()
	m := mux.NewRouter()

	go log.Fatal(http.ListenAndServe(":8081", api.Init()))

	// soon main page
	log.Fatal(http.ListenAndServe(":8080", m))
}
