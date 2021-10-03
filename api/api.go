package api

import (
	"github.com/arcoz0308/arcoz0308.tech/api/arcpaste"
	"github.com/arcoz0308/arcoz0308.tech/api/discord"
	"github.com/arcoz0308/arcoz0308.tech/api/minecraft"
	"github.com/arcoz0308/arcoz0308.tech/middlewares"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	_, err := utils.Cron.AddFunc("* */5 * * * *", func() {
		minecraft.ClearCache()
		middlewares.ClearRateLimitCache()
	})
	if err != nil {
		panic(err)
	}
	api := mux.NewRouter()
	discord.Init(api.PathPrefix("/discord").Subrouter())
	api.HandleFunc("/mcbe/{server}", minecraft.QueryMCBE).Methods("GET")
	api.HandleFunc("/mcpe/{server}", minecraft.QueryMCBE).Methods("GET")
	api.HandleFunc("/arcpaste/{key}", arcpaste.GetPaste).Methods("GET")
	api.Use(middlewares.LogAPIRequest)
	api.Use(middlewares.CheckGlobalAPIRateLimit)
	return api
}
