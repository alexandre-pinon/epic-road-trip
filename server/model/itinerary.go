package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport string

const (
	Airplane Transport = "AIRPLANE"
	Bus      Transport = "BUS"
	Subway   Transport = "SUBWAY"
	Train    Transport = "TRAIN"
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
	Price          float64            `json:"price"`
	Steps          []ItineraryStep    `json:"steps,omitempty" bson:"steps,omitempty"`
}

type ItineraryStep struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Type           Transport          `json:"type"`
	Departure      Station            `json:"departure"`
	Arrival        Station            `json:"arrival"`
	Duration       time.Duration      `json:"-"`
	DurationString string             `json:"duration" bson:"-"`
	Startdate      time.Time          `json:"startdate"`
	Enddate        time.Time          `json:"enddate"`
}
