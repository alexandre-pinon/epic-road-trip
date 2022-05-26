package model

type User struct {
	Firstname      string      `json:"firstname" binding:"required,min=2,max=50"`
	Lastname       string      `json:"lastname" binding:"required,min=2,max=50"`
	Email          string      `json:"email" binding:"required,email"`
	HashedPassword string      `json:"-"`
	Phone          string      `json:"phone" binding:"required,e164,len=12"`
	Trips          []*RoadTrip `json:"roadTrip"`
}
type UserFormData struct {
	User
	Password string `json:"password" binding:"required,min=8,max=100"`
}
