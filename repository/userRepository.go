package repository

import (
	"context"
	"github.com/Tutor2Tutee/T2T-GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userCollection struct {
	*mongo.Collection
}

var UserCollection userCollection

func (c userCollection) Start() {
	UserCollection.Collection = Resource.DB.Collection("user")
}

func (c userCollection) FindUserByID(UserID primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := c.FindOne(context.TODO(), bson.D{{"_id", UserID}}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c userCollection) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := c.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c userCollection) InsertUser(newUser *models.User) (*mongo.InsertOneResult, error) {
	result, err := c.InsertOne(context.TODO(), newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c userCollection) UpdateUserByID(userID primitive.ObjectID, updateUser models.User) (*mongo.UpdateResult, error) {
	result, err := c.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.D{{"$set", updateUser}},
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c userCollection) DeleteUserByID(userID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := c.DeleteOne(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		return nil, err
	}
	return result, err
}
