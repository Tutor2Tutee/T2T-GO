package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthenticationMiddleware(c *gin.Context) {
	if c.Request.Header["Token"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token found in the headers"})
		c.Abort()
		return
	}

	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token found in the headers"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := fmt.Sprintf("%v", claims["UserID"])
		email := fmt.Sprintf("%v", claims["Email"])
		nickname := fmt.Sprintf("%v", claims["Nickname"])
		c.Request.Header.Set("UserID", userID)
		c.Request.Header.Set("UserEmail", email)
		c.Request.Header.Set("UserNickname", nickname)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "No token found in the headers"})
	c.Abort()
	return
}
