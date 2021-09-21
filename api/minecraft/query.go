package minecraft

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandertv/gophertunnel/query"
	"io"
	"net/http"
	"strings"
)

func Query(w http.ResponseWriter, r *http.Request) {
	serv := mux.Vars(r)["server"]
	if len(strings.Split(serv, ":")) == 1 {
		serv = serv + ":25565"
	}
	infos, err := query.Do(serv)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, fmt.Sprintf("{\"error\":\"%s\"}", err.Error()))
		return
	}
	j, _ := json.Marshal(infos)
	io.WriteString(w, string(j))
}
func QueryMCBE(w http.ResponseWriter, r *http.Request) {
	serv := mux.Vars(r)["server"]
	if len(strings.Split(serv, ":")) == 1 {
		serv = serv + ":19132"
	}
	w.Header().Set("Content-Type", "application/json")
	infos, err := query.Do(serv)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, fmt.Sprintf("{\"error\":\"%s\"}", err.Error()))
		return
	}
	w.WriteHeader(200)
	j, _ := json.Marshal(infos)
	io.WriteString(w, string(j))
}
