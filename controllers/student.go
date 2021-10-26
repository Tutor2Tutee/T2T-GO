package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Student{Username: "test student", Password: "123"})
}
