package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	SecretKey := "-----BEGIN CERTIFICATE-----\n" +
		os.Getenv("KEY_KEYCLOAK") +
		"\n-----END CERTIFICATE-----"

	reqToken := c.GetHeader("Authorization")
	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		fmt.Println("Erro1")
		fmt.Println(er)
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token is valid")
	}

}
