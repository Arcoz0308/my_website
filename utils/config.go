package utils

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Config config

type config struct {
	Discord  discordConfig
	Database databaseConfig
}
type discordConfig struct {
	Token string
}
type databaseConfig struct {
	Url string
}

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
	}
}
