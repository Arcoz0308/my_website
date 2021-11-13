package arcpaste

import (
	"database/sql"
	"github.com/arcoz0308/arcoz0308.tech/utils/arcpaste"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{([A-Z]|[a-z])\\w+}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "arcpaste/static/index.html")
	})
	r.HandleFunc("/asset/{file}", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "arcpaste/static/asset/"+mux.Vars(request)["file"])
	})

	r.HandleFunc("/raw/{key}", getRaw)
	return r
}
func getRaw(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	data, err := arcpaste.Data(key)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "404 : not found")
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "500 : internal server error")
		return
	}
	if data.Password != "" {
		password := r.URL.Query().Get("password")
		if password == "" {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "this paste need a password")
			return
		}
		if data.Password != password {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "invalid password")
			return
		}
	}
	io.WriteString(w, data.Raw)
}
