package arcpaste

import (
	"encoding/json"
	"github.com/arcoz0308/arcoz0308.tech/utils/api"
	"github.com/arcoz0308/arcoz0308.tech/utils/arcpaste"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func GetPaste(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	d, err := arcpaste.Data(key)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			api.Error(w, http.StatusNotFound, "paste don't exist")
			return
		}
		log.Print(err)
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if d.Password != "" {
		p := r.URL.Query().Get("password")
		if p == "" {
			api.Error(w, http.StatusForbidden, "this paste are lock with a password")
			return
		}
		if p != d.Password {
			api.Error(w, http.StatusForbidden, "password is invalid")
			return
		}
	}
	b, err := json.Marshal(struct {
		Language string `json:"language"`
		Content  string `json:"content"`
	}{
		d.Language,
		d.Raw,
	})
	if err != nil {
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
