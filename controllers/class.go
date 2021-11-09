package controllers

import (
	"net/http"
	"time"

	"github.com/Tutor2Tutee/T2T-GO/helpers"
	"github.com/Tutor2Tutee/T2T-GO/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
}

func (CC ClassController) GetAll(c *gin.Context) {
	classes, err := repository.ClassCollection.FindAllClasses()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while querying FindAllClasses",
		})
	}

	c.JSON(http.StatusOK, classes)
}

func (CC ClassController) GetOne(c *gin.Context) {
	classID, _ := primitive.ObjectIDFromHex(c.Param("cid"))

	class, err := repository.ClassCollection.FindOneById(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (CC ClassController) Create(c *gin.Context) {
	var class models.Class
	err := c.BindJSON(&class)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	class.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	class.Listener = []models.User{}

	// Find user with user Email
	user, err := repository.UserCollection.FindUserByEmail(class.Teacher.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// JWT Verification
	isAuthenticated, _ := helpers.VerifyUserAuthUsingJWT(c, user.ID.Hex())

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
		return
	}

	result, err := repository.ClassCollection.InsertClass(class)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"createdId": result.InsertedID,
	})
}
