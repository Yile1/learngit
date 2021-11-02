package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"techtrainingcamp-AppUpgrade/common"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	common.RedisInit()

	r := gin.Default()
	customizerouter(r)
	r.Run()
}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}