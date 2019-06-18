package main

import (
	"go-jwt/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,World!")
	})

	r.POST("/auth", auth.Login)
	r.Run()
}
