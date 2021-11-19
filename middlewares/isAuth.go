package middlewares

import (
	"fmt"
	"net/http"
	"os"
	s "strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuth(c *gin.Context) {
	token, err := jwt.Parse(s.Split(c.GetHeader("Authorization"), " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("token not valid")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_S")), nil
	})

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please try to login ", "error": err.Error()})
		return
	}
	if err != nil || !token.Valid {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please try to login ", "error": err.Error()})
		return
	}
	c.Next()
}
