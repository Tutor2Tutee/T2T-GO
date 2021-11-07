package repository

import (
	"context"
	"github.com/Tutor2Tutee/T2T-GO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type classCollection struct {
	*mongo.Collection
}

var ClassCollection classCollection

func (c classCollection) Start() {
	ClassCollection.Collection = Resource.DB.Collection("class")
}

// FindAllClasses will return classes in slice
func (c classCollection) FindAllClasses() ([]models.Class, error) {
	cur, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var classes = make([]models.Class, 0)

	for cur.Next(context.TODO()) {
		var class models.Class
		if err := cur.Decode(&class); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}

// FindOneById will return one class which matches given id
func (c classCollection) FindOneById(id primitive.ObjectID) (models.Class, error) {
	result := c.FindOne(
		context.TODO(),
		bson.D{{
			"id", id,
		}},
	)

	var class models.Class
	if err := result.Decode(&class); err != nil {
		return class, err
	}

	return class, nil
}

// InsertClass will insert a given class to Database
func (c classCollection) InsertClass(class models.Class) (*mongo.InsertOneResult, error) {
	result, err := c.InsertOne(context.TODO(), class)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//func (c classCollection) UpdateClass(classID primitive.ObjectID, )  {
//
//}
