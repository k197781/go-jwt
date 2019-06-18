package auth

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
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
