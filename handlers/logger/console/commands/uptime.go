package commands

import (
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"time"
)

type Uptime struct{}

func (*Uptime) Run(args []string) {
	logger.AppInfof("command:uptime", "uptime : %s", time.Since(utils.Start).String())
}
