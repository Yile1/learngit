package main

import (
	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/controller"
)

func customizerouter(r *gin.Engine) {
	r.POST("/ping", controller.Pong)
	r.POST("/judge", controller.Hit)
	r.POST("/addRule", controller.AddRule)
	r.POST("/deleteRule", controller.DeleteRule)
	r.POST("/disableRule", controller.DisableRule)
	r.POST("/enableRule", controller.EnableRule)
}
