package utils

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Config config

type config struct {
	Database struct {
		Dsn string `json:"dsn"`
	} `json:"database"`
	Discord struct {
		Token string `json:"token"`
	} `json:"discord"`
	Host struct {
		Api struct {
			Host      string `json:"host"`
			Localport int    `json:"localport"`
		} `json:"api"`
		Arcpaste struct {
			Host      string `json:"host"`
			Localport int    `json:"localport"`
		} `json:"arcpaste"`
	} `json:"host"`
}

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
	}
}
