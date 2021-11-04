package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	refreshToken := c.Request.Header.Get("Refresh_token")
	to, err := helpers.VerifyToken(refreshToken)
	if err != nil {
		// in case token is expired
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	if claims, ok := to.Claims.(jwt.MapClaims); ok && to.Valid {
		c.JSON(http.StatusAccepted, claims)
	}

}
