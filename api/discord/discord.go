package discord

import (
	"github.com/arcoz0308/arcoz0308.tech/api/discord/user"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.HandleFunc("/users/{id}", user.Info)
}
