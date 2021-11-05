package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)


type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `json:"email" validate:"email"`
	Password  string             `json:"password" validate:"min=8"`
	Nickname  string             `json:"nickname" validate:"required"`
	Birth     string             `json:"birth" validate:"required"`
	Point     int                `json:"point"`
  Created_At time.Time         `json:"created_at"`
	UserType  string             `json:"user_type" validate:"required"`
	Listening []Class            `json:"listening"`
	Teaching  []Class            `json:"teaching"`
}