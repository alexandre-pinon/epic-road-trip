package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport string

const (
	Airplane Transport = "AIRPLANE"
	Train    Transport = "TRAIN"
)

type Itinerary struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type      Transport          `json:"type"`
	Departure Station            `json:"departure"`
	Arrival   Station            `json:"arrival"`
	Startdate time.Time          `json:"startdate"`
	Enddate   time.Time          `json:"enddate"`
	Price     float64            `json:"price"`
}
