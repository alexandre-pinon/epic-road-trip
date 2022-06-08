package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Activity  struct {
	ID		primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name	string `json:"name" binding:"required,min=2,max=50"`
}