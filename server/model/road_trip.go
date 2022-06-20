package model

import "time"

type RoadTrip struct {
	Name      	string    `json:"name"`
	Startdate 	time.Time `json:"startdate"`
	Enddate   	time.Time `json:"enddate"`
	Locations 	[]int     `json:"locations"`
}
