package user

import (
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/"+id, nil)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "{\"error\":\"internal server error\"}")
		return
	}
	req.Header.Add("Authorization", "Bot "+utils.Config.Discord.Token)
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "{\"error\":\"internal server error\"}")
		return
	}
	if resp.StatusCode == 403 {
		log.Print("discord token are invalid")
		w.WriteHeader(500)
		io.WriteString(w, "{\"error\":\"internal server error\"}")
		return
	}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "{\"error\":\"internal server error\"}")
			return
		}
		io.WriteString(w, string(body))
		return
	}
	if resp.StatusCode == 404 {
		w.WriteHeader(404)
		io.WriteString(w, "{\"error\":\"user not found\"}")
		return
	}
	if resp.StatusCode == 429 {
		w.WriteHeader(500)
		io.WriteString(w, "{\"error\":\" the server are rate limit from discord\"}")
		return
	}
	io.WriteString(w, "{\"error\":\"unknown error\"}")
}
