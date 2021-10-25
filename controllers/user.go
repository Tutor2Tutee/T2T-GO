package controllers

import (
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var newUser models.User

	// Validate Upcoming Data

	// Get user post body
	c.BindJSON(&newUser)

	// Check User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Login successfully", "userDetails": newUser})
}

func RegisterUser(c *gin.Context) {
	var newUser models.User

	// Validate Upcoming Data

	// Get user post body
	c.BindJSON(&newUser)

	// Store User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "userDetails": newUser})
}

func GetUserByID(c *gin.Context) {
	var newUser models.User
	userID := c.Param("userId")
	// Validate Upcoming Data

	// Get user post body
	c.BindJSON(&newUser)

	// Check User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Found user successfully", "userDetails": newUser})
}

func UpdateUserByID(c *gin.Context) {
	var newUser models.User
	userID := c.Param("userId")
	// Validate Upcoming Data

	// Get user post body
	c.BindJSON(&newUser)

	// Check User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Updated user successfully", "userDetails": newUser})
}

func DeleteUserByID(c *gin.Context) {
	var newUser models.User
	userID := c.Param("userId")

	// Validate Upcoming Data

	// Get user post body
	c.BindJSON(&newUser)

	// Check User in Database
	// *****

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Deleted user successfully", "userDetails": newUser})
}
