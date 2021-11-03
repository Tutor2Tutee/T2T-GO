package controllers

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	ClassCollection *mongo.Collection
	UserCollection  *mongo.Collection
	QuizCollection  *mongo.Collection
}

var Collections Collection

func Start(database *db.Resource) {
	Resource := database
	Collections.ClassCollection = Resource.DB.Collection("class")
	Collections.UserCollection = Resource.DB.Collection("user")
	Collections.QuizCollection = Resource.DB.Collection("quiz")
}
