package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateQuiz(c *gin.Context) {
	var newQuiz models.Quiz

	// Get Request Data
	err := c.BindJSON(&newQuiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validation
	validate := validator.New()
	if err := validate.Struct(&newQuiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create Quiz in DB
	result, err := Collections.QuizCollection.InsertOne(context.Background(), newQuiz)

	fmt.Println("Result", result)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Quiz created successfully", "quiz": newQuiz})
}

func GetAllQuiz(c *gin.Context) {
	var result []models.Quiz

	r, err := Collections.QuizCollection.Find(context.Background(), bson.D{})
	r.All(context.Background(), &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, gin.H{"quiz": result})

}

func GetQuizByID(c *gin.Context) {
	quizID := c.Param("quizID")

	objectID, err := primitive.ObjectIDFromHex(quizID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid quiz id",
		})
		return
	}

	// find
	var quiz models.Quiz
	error := Collections.QuizCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&quiz)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No quiz exists with provided email.",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Found quiz successfully", "quiz": quiz})
}

func GetQuizByCreatorID(c *gin.Context) {
	creatorID := c.Param("creatorID")

	objectID, err := primitive.ObjectIDFromHex(creatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid creator id",
		})
		return
	}

	// find
	var quizzes []models.Quiz
	r, error := Collections.QuizCollection.Find(context.Background(), bson.M{"creator": objectID})
	r.All(context.Background(), &quizzes)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No quizzes exists with provided creator ID.",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusCreated, gin.H{"message": "Found quizzes successfully", "quiz": quizzes})
}
