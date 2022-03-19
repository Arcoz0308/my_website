package redis

import (
	"context"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Username: config.Redis.User,
		Password: config.Redis.Passwd,
		DB:       0,
	})
	err := Client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}
