package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Station struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	City    string             `json:"city"`
	Country string             `json:"country"`
}
