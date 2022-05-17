package entity

type Activity struct {
	ID			uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name		string `json:"name"`
	Description	string `json:"description"`
 }