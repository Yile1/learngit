package main

import (
	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/service"
)

func customizerouter(r *gin.Engine) {
	r.POST("/ping", service.Pong)
	r.POST("/judge", service.Hit)
}
