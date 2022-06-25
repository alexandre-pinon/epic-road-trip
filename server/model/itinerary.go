package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport string

const (
	Air    Transport = "AIR"
	Ground Transport = "GROUND"
)

type Itinerary struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Type           Transport          `json:"type"`
	Departure      Station            `json:"departure"`
	Arrival        Station            `json:"arrival"`
	Duration       time.Duration      `json:"-"`
	DurationString string             `json:"duration" bson:"-"`
	Startdate      time.Time          `json:"startdate"`
	Enddate        time.Time          `json:"enddate"`
	Price          float64            `json:"price,omitempty" bson:"omitempty"`
	Steps          []ItineraryStep    `json:"steps,omitempty" bson:"omitempty"`
}

type ItineraryStep struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Type           string             `json:"type"`
	Departure      string             `json:"departure"`
	Arrival        string             `json:"arrival"`
	Duration       time.Duration      `json:"-"`
	DurationString string             `json:"duration" bson:"-"`
	Startdate      time.Time          `json:"startdate"`
	Enddate        time.Time          `json:"enddate"`
}
