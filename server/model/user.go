package model

type User struct {
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Phone     string      `json:"phone"`
	Trips     []*RoadTrip `json:"roadTrip"`
}
