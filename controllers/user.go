package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	var User models.User

	err := c.BindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check User in Database
	var foundUser *models.User
	error := Collections.UserCollection.FindOne(context.Background(), bson.D{{"email", User.Email}}).Decode(&foundUser)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No user exists with provided email.",
		})
		return
	}

	isPassSame := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(User.Password))

	fmt.Println("Is Password Same", isPassSame)
	if isPassSame != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Login successfully", "userDetails": foundUser})
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

	// Check if User already exits in Database
	var existingUser *models.User
	doesUserExist := Collections.UserCollection.FindOne(context.Background(), bson.D{{"email", newUser.Email}}).Decode(&existingUser)

	if doesUserExist == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User already exists with provided email!"})
		return
	}

	// Hashing the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser.Password = string(bytes)

	// Store User in Database
	result, err := Collections.UserCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userDetails": result})
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
