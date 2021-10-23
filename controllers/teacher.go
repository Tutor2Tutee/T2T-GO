package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
)

func GetTeacher(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Teacher{Username: "test teacher", Password: "123"})
}
