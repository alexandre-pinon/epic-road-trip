package entity

/* import "time"
 */
type Trip struct {
/* 	ID			uint64 `gorm:"primary_key;auto_increment" json:"id"`
	User_id		uint64 `json:"-"`
	User 		User `json:"user" binding:"required" gorm:"foreignkey:UserID"`
 */	Name		string `json:"name"`
/* 	startDate	time.Time `json:"start" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	endDate		time.Time `json:"end" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
 */ }