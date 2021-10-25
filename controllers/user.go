package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.User{Email: "Hello"})
}
