package common

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"io"
	"os"
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

	RDCPipeline := RedisClient.Pipeline()

	fi, err := os.Open(viper.GetString("file.whiteList"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
		RDCPipeline.SAdd("WhiteList", string(a))
	}
	RDCPipeline.Exec()
}

func GetRdc() *redis.Client {
	return RedisClient
}

func RedisTest() {
	RDCPipeline := RedisClient.Pipeline()

	fi, err := os.Open(viper.GetString("file.whiteList"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
		RDCPipeline.SAdd("WhiteListTest", string(a))
	}
	RDCPipeline.Exec()

	fmt.Println(RedisClient.SMembers("WhiteListTest").Val())
	fmt.Println(RedisClient.SIsMember("WhiteListTest", "789").Val())
}