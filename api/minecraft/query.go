package minecraft

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandertv/gophertunnel/query"
	"io"
	"net/http"
	"strings"
	"time"
)

var cache = map[string]serverInfos{}

type serverInfos struct {
	Infos map[string]string
	Time  time.Time
}

func Query(w http.ResponseWriter, r *http.Request) {
	serv := mux.Vars(r)["server"]
	if len(strings.Split(serv, ":")) == 1 {
		serv = serv + ":25565"
	}
	if i, ok := cache[serv]; ok {
		if i.Time.Before(time.Now()) {
			j, _ := json.Marshal(i)
			io.WriteString(w, string(j))
			return
		}
	}

	infos, err := query.Do(serv)
	cache[serv] = serverInfos{
		infos,
		time.Now().Add(time.Second * 20),
	}
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
	if i, ok := cache[serv]; ok {
		if !i.Time.Before(time.Now()) {
			j, _ := json.Marshal(i)
			io.WriteString(w, string(j))
			return
		}
	}

	infos, err := query.Do(serv)
	cache[serv] = serverInfos{
		infos,
		time.Now().Add(time.Second * 20),
	}
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, fmt.Sprintf("{\"error\":\"%s\"}", err.Error()))
		return
	}
	w.WriteHeader(200)
	j, _ := json.Marshal(infos)
	io.WriteString(w, string(j))
}
func ClearCache() {
	for ip, i := range cache {
		if i.Time.Before(time.Now()) {
			delete(cache, ip)
		}
	}
}
