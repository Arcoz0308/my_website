package main

import (
	"github.com/arcoz0308/arcoz0308.tech/api"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

func main() {
	utils.LoadConfig()
	database.Init()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	m := mux.NewRouter()

	m.HandleFunc("/{([A-Z])|[a-z]\\w+}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./arcpaste/index.html")
	})

	m.HandleFunc("/asset/{file}", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./arcpaste/asset/"+mux.Vars(request)["file"])
	})

	go func() {
		m := mux.NewRouter()
		m.HandleFunc("/{([A-Z])|[a-z]\\w+}", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./arcpaste/index.html")
		})

		m.HandleFunc("/asset/{file}", func(writer http.ResponseWriter, request *http.Request) {
			http.ServeFile(writer, request, "./arcpaste/asset/"+mux.Vars(request)["file"])
		})
		log.Fatal(http.ListenAndServe(":8080", m))
		wg.Done()
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8081", api.Init()))
		wg.Done()
	}()
	wg.Wait()
}
