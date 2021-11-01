package main

import (
	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/controller"
	"techtrainingcamp-AppUpgrade/service"
)

func customizerouter(r *gin.Engine) {
	r.POST("/ping", service.Pong)
	r.POST("/judge", controller.Hit)
}
