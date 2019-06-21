package auth

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func Login(c *gin.Context) {
	username := "test"
	password := "test"

	if username == "test" && password == "test" {
		// create token
		token := jwt.New(jwt.SigningMethodHS256)

		// set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "test"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// create signature
		tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, tokenString)
	}
}

var IsSignin = func(c *gin.Context) {
	// check token
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SIGNINGKEY"))
		return b, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("token error")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
