package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/li-jin-gou/hz_demo/conf"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
