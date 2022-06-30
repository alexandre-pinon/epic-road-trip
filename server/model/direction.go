package model

import "time"

type DirectionsFormData struct {
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DepartureTime time.Time `json:"departureTime"`
}

type GoogleDirectionsResponse struct {
	GeocodedWaypoints []GeocodedWaypoints `json:"geocoded_waypoints"`
	Routes            []Routes            `json:"routes"`
	Status            string              `json:"status"`
}

type GeocodedWaypoints struct {
	GeocoderStatus string   `json:"geocoder_status"`
	PlaceID        string   `json:"place_id"`
	Types          []string `json:"types"`
}

type GoogleTime struct {
	Text     string `json:"text"`
	TimeZone string `json:"time_zone"`
	Value    int    `json:"value"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Polyline struct {
	Points string `json:"points"`
}

type Stop struct {
	Location Location `json:"location"`
	Name     string   `json:"name"`
}

type Agencies struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	URL   string `json:"url"`
}

type Vehicle struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Line struct {
	Agencies  []Agencies `json:"agencies"`
	Color     string     `json:"color"`
	Name      string     `json:"name"`
	ShortName string     `json:"short_name"`
	TextColor string     `json:"text_color"`
	Vehicle   Vehicle    `json:"vehicle"`
}

type TransitDetails struct {
	ArrivalStop   Stop       `json:"arrival_stop"`
	ArrivalTime   GoogleTime `json:"arrival_time"`
	DepartureStop Stop       `json:"departure_stop"`
	DepartureTime GoogleTime `json:"departure_time"`
	Headsign      string     `json:"headsign"`
	Line          Line       `json:"line"`
	NumStops      int        `json:"num_stops"`
}

type GoogleDestinationSubStep struct {
	Distance         Distance `json:"distance"`
	Duration         Duration `json:"duration"`
	EndLocation      Location `json:"end_location"`
	HTMLInstructions string   `json:"html_instructions"`
	Polyline         Polyline `json:"polyline"`
	StartLocation    Location `json:"start_location"`
	TravelMode       string   `json:"travel_mode"`
}

type GoogleDestinationStep struct {
	Distance         Distance                   `json:"distance"`
	Duration         Duration                   `json:"duration"`
	EndLocation      Location                   `json:"end_location"`
	HTMLInstructions string                     `json:"html_instructions"`
	Polyline         Polyline                   `json:"polyline"`
	StartLocation    Location                   `json:"start_location"`
	TransitDetails   TransitDetails             `json:"transit_details,omitempty"`
	TravelMode       string                     `json:"travel_mode"`
	SubSteps         []GoogleDestinationSubStep `json:"steps,omitempty"`
}

type Legs struct {
	ArrivalTime       GoogleTime              `json:"arrival_time"`
	DepartureTime     GoogleTime              `json:"departure_time"`
	Distance          Distance                `json:"distance"`
	Duration          Duration                `json:"duration"`
	EndAddress        string                  `json:"end_address"`
	EndLocation       Location                `json:"end_location"`
	StartAddress      string                  `json:"start_address"`
	StartLocation     Location                `json:"start_location"`
	Steps             []GoogleDestinationStep `json:"steps"`
	TrafficSpeedEntry []interface{}           `json:"traffic_speed_entry"`
	ViaWaypoint       []interface{}           `json:"via_waypoint"`
}

type Routes struct {
	Bounds           Bounds        `json:"bounds"`
	Copyrights       string        `json:"copyrights"`
	Legs             []Legs        `json:"legs"`
	OverviewPolyline Polyline      `json:"overview_polyline"`
	Summary          string        `json:"summary"`
	Warnings         []string      `json:"warnings"`
	WaypointOrder    []interface{} `json:"waypoint_order"`
}
