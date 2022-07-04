package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Roadtrip struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name"`
	Startdate   time.Time            `json:"startdate"`
	Enddate     time.Time            `json:"enddate"`
	TripSteps   []primitive.ObjectID `json:"tripSteps"`
	Itineraries []primitive.ObjectID `json:"itineraries"`
}

type TripStep struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	City      string             `json:"city"`
	Startdate time.Time          `json:"startdate"`
	Enddate   time.Time          `json:"enddate"`
	Travel    Itinerary          `json:"travel"`
	Enjoy     Enjoy              `json:"enjoy"`
	Sleep     Sleep              `json:"sleep"`
	Eat       Eat                `json:"eat"`
	Drink     Drink              `json:"drink"`
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
