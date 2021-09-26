package user

import (
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"github.com/arcoz0308/arcoz0308.tech/utils/api"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/**
@api {GET} /discord/users/:id discord user info
@apiDescription get infos about a discord user without a discord token

*/

func Info(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/"+id, nil)
	if err != nil {
		log.Print(err)
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	req.Header.Add("Authorization", "Bot "+utils.Config.Discord.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if resp.StatusCode == 403 {
		log.Print("discord token are invalid")
		api.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			api.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
		io.WriteString(w, string(body))
		return
	}
	if resp.StatusCode == 404 {
		api.Error(w, http.StatusNotFound, "user not found")
		return
	}
	if resp.StatusCode == 429 {
		api.Error(w, http.StatusInternalServerError, "the api are by rate limit from discord")
		return
	}
	api.Error(w, http.StatusInternalServerError, "internal server error")
}
