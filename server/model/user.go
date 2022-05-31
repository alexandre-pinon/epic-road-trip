package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname      string             `json:"firstname" binding:"required,min=2,max=50"`
	Lastname       string             `json:"lastname" binding:"required,min=2,max=50"`
	Email          string             `json:"email" binding:"required,email"`
	HashedPassword string             `json:"-" bson:"hashedpassword,omitempty"`
	Phone          string             `json:"phone" binding:"required,e164,len=12"`
	Trips          []*RoadTrip        `json:"roadTrip,omitempty" bson:"trips,omitempty"`
}
type UserFormData struct {
	User
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
