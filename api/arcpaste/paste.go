package arcpaste

import (
	"database/sql"
	"encoding/json"
	"github.com/arcoz0308/arcoz0308.tech/utils/api"
	"github.com/arcoz0308/arcoz0308.tech/utils/arcpaste"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type dataStruct struct {
	Raw      string  `json:"raw"`
	Password *string `json:"password"`
	Expire   *int    `json:"expire"`
	Language *string `json:"language"`
}

func GetPaste(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	d, err := arcpaste.Data(key)
	if err != nil {
		if err == sql.ErrNoRows {
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
func PostPaste(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	var c dataStruct
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Print(err)
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

}

func generateKey() string {
	rand.Seed(time.Now().UnixMilli())
	c := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	n := rand.Intn(4) + 8
	s := make([]byte, n)
	for {
		for i := range s {
			s[i] = c[rand.Intn(len(c))]
		}
		_, err := arcpaste.Data(string(s))
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return string(s)
			} else {
				panic(err)
			}
		}
	}
}
