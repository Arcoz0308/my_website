package utils

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Config config

type config struct {
	Discord discordConfig
}
type discordConfig struct {
	Token string
}

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
	}
}
