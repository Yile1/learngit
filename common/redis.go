package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"), //@hld windows上的redis好像没密码,需要注释此行才能运行
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}

func GetRdc() *redis.Client {
	return RedisClient
}
