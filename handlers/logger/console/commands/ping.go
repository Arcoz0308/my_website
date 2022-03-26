package commands

import (
	"github.com/arcoz0308/arcoz0308.tech/handlers/database"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/handlers/redis"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"sync"
)

type Ping struct{}

func (*Ping) Run(args []string) {

	// async
	wg := sync.WaitGroup{}
	wg.Add(2)

	// database ping
	go func() {
		ping, err := database.Ping()
		if err != nil {
			logger.AppErrorf("command:ping", "failed to ping database, error : %s", err.Error())
		} else {
			logger.AppInfof("command:ping", "database ping : %s", utils.MsWith2Decimal(ping))
		}
		wg.Done()
	}()

	// redis ping
	go func() {
		ping, err := redis.Ping()
		if err != nil {
			logger.AppErrorf("command:ping", "failed to ping redis, error : %s", err.Error())
		} else {
			logger.AppInfof("command:ping", "redis ping : %s", utils.MsWith2Decimal(ping))
		}
		wg.Done()
	}()
	wg.Wait()
}
