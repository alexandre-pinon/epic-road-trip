package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoadTrip struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name"`
	Startdate   time.Time            `json:"startdate"`
	Enddate     time.Time            `json:"enddate"`
	TripSteps   []primitive.ObjectID `json:"tripSteps"`
	Itineraries []primitive.ObjectID `json:"itineraries"`
}
