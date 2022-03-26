package config

import (
	"github.com/BurntSushi/toml"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"time"
)

var (
	Cert     *cert
	Database *database
	Redis    *redis
	Discord  *discord
	Global   *global
)

type config struct {
	Cert     cert     `json:"cert"`
	Database database `json:"database"`
	Redis    redis    `json:"redis"`
	Discord  discord  `json:"discord"`
	Global   global   `json:"global"`
}

type cert struct {
	Addrs []string `json:"addrs"`
	Dir   string   `json:"dir"`
	Email string   `json:"email"`
}

type database struct {
	Addr   string `json:"addr"`
	DbName string `json:"db_name"`
	Passwd string `json:"passwd"`
	User   string `json:"user"`
}

type redis struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

type discord struct {
	Token []string `json:"token"`
}
type global struct {
	Host string `json:"host"`
}

func Init() {
	t := time.Now()
	var c config
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		panic(err)
	}
	Cert = &c.Cert
	Database = &c.Database
	Redis = &c.Redis
	Discord = &c.Discord
	Global = &c.Global
	logger.Infof("loaded config in %s", utils.MsWith2Decimal(time.Since(t)))
}
