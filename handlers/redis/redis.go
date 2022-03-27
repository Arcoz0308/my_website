package redis

import (
	"context"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/go-redis/redis/v8"
	"time"
)

var Client *redis.Client

var (
	PrefixLongCookieToken = "a"
)

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Username: config.Redis.User,
		Password: config.Redis.Passwd,
		DB:       0,
	})
	_, err := Ping()
	if err != nil {
		logger.AppFatal(true, "redis", err)
	}
}

func Ping() (time.Duration, error) {
	t := time.Now()
	err := Client.Ping(context.Background()).Err()
	if err != nil {
		return -1, err
	}

	return time.Since(t), nil
}
