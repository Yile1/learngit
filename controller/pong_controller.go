package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "PONG"})
}

func Form(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}
