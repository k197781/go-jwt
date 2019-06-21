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

	// check token
	r.POST("/auth", auth.Login)

	auth := r.Group("/", auth.IsSignin)

	auth.GET("/private", func(c *gin.Context) {
		c.String(200, "success!")
	})
	r.Run()
}
