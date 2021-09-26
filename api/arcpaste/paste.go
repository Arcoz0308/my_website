package arcpaste

import (
	"encoding/json"
	"github.com/arcoz0308/arcoz0308.tech/utils/arcpaste"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func GetPaste(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	key := mux.Vars(r)["key"]
	d, err := arcpaste.Data(key)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			b, _ := json.Marshal("{\"error\":\"element not found\"}")
			w.Write(b)
			return
		}
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal("{\"error\":\"internal server error\"}")
		w.Write(b)
		return
	}
	b, err := json.Marshal(struct {
		Language string `json:"language"`
		Content  string `json:"content"`
	}{
		d.Language,
		d.Raw,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal("{\"error\":\"internal server error\"}")
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
