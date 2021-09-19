package api

import (
	"github.com/arcoz0308/arcoz0308.tech/api/discord"
	"github.com/arcoz0308/arcoz0308.tech/api/minecraft"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	api := mux.NewRouter()
	discord.Init(api.PathPrefix("/discord").Subrouter())
	api.HandleFunc("/mc/{server}", minecraft.Query)
	api.HandleFunc("/minecraft/{server}", minecraft.Query)
	api.HandleFunc("/mcbe/{server}", minecraft.QueryMCBE)
	api.HandleFunc("/mcpe/{server}", minecraft.QueryMCBE)
	return api
}
