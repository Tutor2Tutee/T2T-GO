package controllers

import (
	"fmt"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	refreshToken := c.Request.Header.Get("Refresh_token")
	if c.Request.Header["Refresh_token"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No refresh token found in the headers"})
		return
	}

	to, err := helpers.VerifyToken(refreshToken)
	if err != nil {
		// in case token is expired
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	claims, _ := to.Claims.(jwt.MapClaims)

	// Verify
	//

	// Return new token
	email := fmt.Sprintf("%v", claims["Email"])
	nickname := fmt.Sprintf("%v", claims["Nickname"])
	userID := fmt.Sprintf("%v", claims["UserID"])

	token := helpers.GenerateAccessToken(email, nickname, userID)

	c.JSON(http.StatusCreated, gin.H{"message": "generated new token successfully", "token": token})

}
