package entity

// Imports
import "time"

// Model of User System
type UserSystem struct {

	// User System Model Attributes
	ID                 uint64    `json:"id" gorm:"primary_key;auto_increment;unique;not null"`
	UserSystemName     string    `json:"user_system_name" binding:"min=3,max=50" gorm:"type:varchar(50);not null"`
	UserSystemEmail    string    `json:"user_system_email" binding:"min=3,max=50" gorm:"type:varchar(50);unique;not null"`
	UserSystemPassword string    `json:"user_system_password" binding:"min=3,max=50" gorm:"type:varchar(50);unique;not null"`
	CreatedAt          time.Time `json:"-" time_format:"2006-01-02" time_utc:"1"`
	IsActive           int       `json:"is_active"`
}
