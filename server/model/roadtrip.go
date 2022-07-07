package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Roadtrip struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Startdate   time.Time            `json:"startdate"`
	Enddate     time.Time            `json:"enddate"`
	TripStepsID []primitive.ObjectID `json:"-" bson:"tripSteps_id"`
	TripSteps   *[]TripStep          `json:"tripSteps,omitempty" bson:"tripSteps,omitempty"`
}

type TripStep struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	City      string             `json:"city" binding:"required"`
	Startdate time.Time          `json:"startdate" binding:"required"`
	Enddate   time.Time          `json:"enddate" binding:"required"`
	Travel    *Itinerary         `json:"travel,omitempty" bson:"travel,omitempty"`
	Enjoy     *[]Enjoy           `json:"enjoy,omitempty"`
	Sleep     *[]Sleep           `json:"sleep,omitempty"`
	Eat       *[]Eat             `json:"eat,omitempty"`
	Drink     *[]Drink           `json:"drink,omitempty"`
}
type Enjoy struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rating   float64 `json:"rating"`
	Vicinity string  `json:"vicinity"`
}
type Sleep struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rating   float64 `json:"rating"`
	Vicinity string  `json:"vicinity"`
}
type Eat struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rating   float64 `json:"rating"`
	Vicinity string  `json:"vicinity"`
}
type Drink struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rating   float64 `json:"rating"`
	Vicinity string  `json:"vicinity"`
}
