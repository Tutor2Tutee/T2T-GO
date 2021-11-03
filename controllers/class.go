package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Tutor2Tutee/T2T-GO/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(c *gin.Context) {
	cur, err := Collections.ClassCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
	}
	var classes = make([]models.Class, 0)

	for cur.Next(context.TODO()) {
		var class models.Class
		err := cur.Decode(&class)
		if err != nil {
			log.Println(err)
		}
		classes = append(classes, class)
	}

	c.JSON(http.StatusOK, classes)
}

func Create(c *gin.Context) {
	var class models.Class
	err := c.BindJSON(&class)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result, err := Collections.ClassCollection.InsertOne(context.TODO(), class)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"createdId": result.InsertedID,
	})
}
