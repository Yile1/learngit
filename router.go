package main

import (
	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/controller"
)

func customizerouter(r *gin.Engine) {
	r.GET("/ping", controller.Pong)
	r.POST("/judge", controller.Hit)
	r.GET("/rule", controller.GetRule)
	r.POST("/rule", controller.AddRule)
	r.DELETE("/rule/:id", controller.DeleteRule)
	r.POST("/disable-rule", controller.DisableRule)
	r.POST("/enable-rule", controller.EnableRule)
}
