package controllers

import (
	"fmt"
	"github.com/Tutor2Tutee/T2T-GO/repository"
	"net/http"
	"time"

	"github.com/Tutor2Tutee/T2T-GO/helpers"
	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	foundUser, err := repository.UserCollection.FindUserByEmail(User.Email)
	if err != nil {
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
	// Generate Token
	token := helpers.GenerateAccessToken(foundUser.Email, foundUser.Nickname, foundUser.ID.Hex())
	refreshToken := helpers.GenerateRefreshToken(foundUser.Email, foundUser.Nickname, foundUser.ID.Hex())

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Login successfully", "userDetails": foundUser, "access_token": token, "refresh_token": refreshToken})
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
	//var existingUser *models.User
	//doesUserExist := repository.UserCollection.FindOne(context.Background(), bson.D{{"email", newUser.Email}}).Decode(&existingUser)
	_, err = repository.UserCollection.FindUserByEmail(newUser.Email)
	if err == nil {
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
	newUser.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newUser.Teaching = []models.Class{}
	newUser.Listening = []models.Class{}

	// Store User in Database
	result, err := repository.UserCollection.InsertUser(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userDetails": result})
}

func GetUserByID(c *gin.Context) {
	// Get userID from params
	userID := c.Param("userId")

	// Check User in Database and return
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// find
	//var user models.User
	//error := repository.UserCollection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)

	user, err := repository.UserCollection.FindUserByID(objectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No user exists with provided id.",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Found user successfully", "user": user})
}

func UpdateUserByID(c *gin.Context) {
	// Get userID from params
	userID := c.Param("userId")

	// Create ObjectID
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Check User in Database
	//var foundResult models.User
	_, err = repository.UserCollection.FindUserByID(objectId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No user exists with provided user id.",
		})
		return
	}

	// Get Upcoming Body Data
	var newUser models.User
	err = c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate upcoming data
	validate := validator.New()
	if err := validate.Struct(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser.Password = string(bytes)

	// Update Data in the database
	result, updateError := repository.UserCollection.UpdateUserByID(objectId, newUser)
	if updateError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": updateError.Error(),
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Updated user successfully", "user": result})
}

func DeleteUserByID(c *gin.Context) {
	// Get userID from params
	userID := c.Param("userId")

	// Create ObjectID
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	res, _ := repository.UserCollection.DeleteUserByID(objectId)
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No user exists with provided user id.",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Deleted user successfully"})
}
