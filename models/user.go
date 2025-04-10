package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"size:30;not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Email     string    `json:"email" gorm:"size:255;uniqueIndex;not null"`
	BirthDay  string    `json:"birthday"`
	Gender    string    `json:"gender" gorm:"size:10"`
	PhotoURL  string    `json:"photo_url" gorm:"size:255"`
	LastLogin time.Time `json:"last_login" gorm:"autoCreateTime:false"`
	Active    bool      `json:"active" gorm:"default:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime:false"`
	Password  string    `json:"-" gorm:"size:255;not null"`
}

// Note: Use this func to override the default table-name
func (User) TableName() string {
	return "users"
}
