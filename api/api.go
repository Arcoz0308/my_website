package api

import (
	"github.com/arcoz0308/arcoz0308.tech/api/arcpaste"
	"github.com/arcoz0308/arcoz0308.tech/api/discord"
	"github.com/arcoz0308/arcoz0308.tech/api/minecraft"
	"github.com/arcoz0308/arcoz0308.tech/middlewares"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"time"
)

func Init() *mux.Router {
	task := gocron.NewScheduler(time.UTC)
	task.Every(5).Minutes().Do(func() {
		minecraft.ClearCache()
		middlewares.ClearRateLimitCache()
	})
	api := mux.NewRouter()
	discord.Init(api.PathPrefix("/discord").Subrouter())
	api.HandleFunc("/mc/{server}", minecraft.Query).Methods("GET")
	api.HandleFunc("/minecraft/{server}", minecraft.Query).Methods("GET")
	api.HandleFunc("/mcbe/{server}", minecraft.QueryMCBE).Methods("GET")
	api.HandleFunc("/mcpe/{server}", minecraft.QueryMCBE).Methods("GET")
	api.HandleFunc("/arcpaste/{key}", arcpaste.GetPaste).Methods("GET")
	api.Use(middlewares.LogAPIRequest)
	api.Use(middlewares.CheckGlobalAPIRateLimit)
	return api
}
