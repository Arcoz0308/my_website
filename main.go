package main

import (
	"github.com/arcoz0308/arcoz0308.tech/api"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/arcoz0308/arcoz0308.tech/utils/arcpaste"
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
	"github.com/arcoz0308/arcoz0308.tech/utils/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	utils.LoadCron()
	logger.Init()
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
		m.HandleFunc("/{([A-Z]|[a-z])\\w+}", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./arcpaste/index.html")
		})
		m.HandleFunc("/asset/{file}", func(writer http.ResponseWriter, request *http.Request) {
			http.ServeFile(writer, request, "./arcpaste/asset/"+mux.Vars(request)["file"])
		})
		m.HandleFunc("/raw/{key}", func(w http.ResponseWriter, r *http.Request) {
			p, err := arcpaste.Data(mux.Vars(r)["key"])
			if err != nil {
				if err == mongo.ErrNoDocuments {
					w.WriteHeader(http.StatusNotFound)
					io.WriteString(w, "404 not found")
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "500 internal server error")
				return
			}
			if p.Password != "" {
				password := r.URL.Query().Get("password")
				if password == "" {
					w.WriteHeader(http.StatusForbidden)
					io.WriteString(w, "this paste need a password")
					return
				}
				if p.Password != password {
					w.WriteHeader(http.StatusForbidden)
					io.WriteString(w, "invalid password")
					return
				}
			}
			io.WriteString(w, p.Raw)
		})
		log.Fatal(http.ListenAndServe(":8080", m))
		wg.Done()
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8081", api.Init()))
		wg.Done()
	}()
	utils.StartCron()
	wg.Wait()
}
