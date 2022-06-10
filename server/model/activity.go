package model

/* import "go.mongodb.org/mongo-driver/bson/primitive" */

type Open struct {
	Open_day		int `json:"open_day""`
	Open_time 	string `json:"open_time""`
}

type Close struct {
	Close_day		int `json:"close_day""`
	Close_time	string `json:"close_time""`
}

type Periods struct {
	Close_periods	[]*Close 	`json:"close," bson:"close,"`
	Open_periods	[]*Open 	`json:"open," bson:"open,"`
}


type Opening_hours struct {
	Open_now bool 		`json:"open_now" binding:"required"`
	Periods  []*Periods	`json:"periods," bson:"periods,"`
}

type Photos struct {
	Height			int
	Width			int
	Photo_reference	string
}

type Activity  struct {
	/* ID		primitive.ObjectID `json:"id," bson:"_id,"` */
	Place_id 			string 			`json:"place_id" binding:"required"`
	Name				string 			`json:"name" binding:"required,min=2,max=50"`
	Opening_hours		[]*Opening_hours `json:"opening_hours," bson:"opening,"`
	Address 			string 			`json:"address" binding:"required`
	Types 				string  		`json:"types" binding:"required`
	User_ratings_total	int 			`json:"user_ratings_total" binding:"required`
	Price_level			int 			`json:"price_level" binding:"required`
	Rating				float64 		`json:"rating" binding:"required`
}