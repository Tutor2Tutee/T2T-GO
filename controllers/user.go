package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginUser(c *gin.Context) {
	var User models.User

	// Validate Upcoming Data
	err := c.BindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Login successfully", "userDetails": User})
}

func RegisterUser(c *gin.Context) {
	var newUser models.User

	// Validate Upcoming Data
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	validate := validator.New()
	if err := validate.Struct(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Store User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userDetails": newUser})
}

func GetUserByID(c *gin.Context) {
	// Get userID from params
	// userID := c.Param("userId")

	// Check User in Database and return
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Found user successfully"})
}

func UpdateUserByID(c *gin.Context) {
	// Get userID from params
	// userID := c.Param("userId")

	// Validate Upcoming Data
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check User in Database and update
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Updated user successfully"})
}

func DeleteUserByID(c *gin.Context) {
	// Get userID from params
	// userID := c.Param("userId")

	// Check User in Database and delete
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Deleted user successfully"})
}
