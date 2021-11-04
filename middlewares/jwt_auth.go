package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthenticationMiddleware(c *gin.Context) {
	if c.Request.Header["Token"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token found in the headers"})
		c.Abort()
		return
	}

	token, err := helpers.VerifyToken(c.Request.Header.Get("Token"))

	if err != nil {
		// in case token is expired
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
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
