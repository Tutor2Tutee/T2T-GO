package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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
