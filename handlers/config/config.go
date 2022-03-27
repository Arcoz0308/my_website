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
	Cert     cert     `toml:"cert"`
	Database database `toml:"database"`
	Redis    redis    `toml:"redis"`
	Discord  discord  `toml:"discord"`
	Global   global   `toml:"global"`
}

type cert struct {
	CertFile string `toml:"cert_file"`
	Key      string `toml:"key"`
}

type database struct {
	Addr   string `toml:"addr"`
	DbName string `toml:"db_name"`
	Passwd string `toml:"passwd"`
	User   string `toml:"user"`
}

type redis struct {
	Addr   string `toml:"addr"`
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
}

type discord struct {
	Token []string `toml:"token"`
}
type global struct {
	Host string `toml:"host"`
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
