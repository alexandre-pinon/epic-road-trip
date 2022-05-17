package entity

type User struct {
/* 	ID			uint64 `gorm:"primary_key;auto_increment" json:"id"`
 */	Lastname	string `json:"Lastname"`
	Firstname	string 	`json:"Firstname"`
}