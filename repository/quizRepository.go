package repository

import (
	"context"
	"github.com/Tutor2Tutee/T2T-GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type quizCollection struct {
	*mongo.Collection
}

var QuizCollection quizCollection

func (c quizCollection) Start() {
	QuizCollection.Collection = Resource.DB.Collection("quiz")
}

func (c quizCollection) Create(newQuiz models.Quiz) (*mongo.InsertOneResult, error) {
	result, err := c.InsertOne(context.TODO(), newQuiz)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c quizCollection) GetAllQuiz() ([]models.Quiz, error) {
	cur, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var quizzes []models.Quiz
	if err := cur.All(context.TODO(), &quizzes); err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (c quizCollection) GetQuizById(ObjectID primitive.ObjectID) (*models.Quiz, error) {
	var quiz models.Quiz
	result := c.FindOne(context.TODO(), bson.D{{"_id", ObjectID}})
	if err := result.Decode(&quiz); err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (c quizCollection) GetQuizByCreatorID(creatorID primitive.ObjectID) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	result, err := c.Find(context.TODO(), bson.D{{"creator", creatorID}})
	if err != nil {
		return nil, err
	}
	if err := result.All(context.TODO(), &quizzes); err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (c quizCollection) DeleteOneById(quizID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := c.DeleteOne(context.TODO(), bson.D{{"_id", quizID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}
