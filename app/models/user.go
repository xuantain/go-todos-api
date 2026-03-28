package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"size:30;not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Email     string    `json:"email" gorm:"size:255;uniqueIndex;not null"`
	BirthDay  string    `json:"birthday"`
	Gender    string    `json:"gender" gorm:"size:10"`
	PhotoURL  string    `json:"photoUrl" gorm:"size:255"`
	LastLogin time.Time `json:"lastLogin" gorm:"autoCreateTime:false"`
	Active    bool      `json:"active" gorm:"default:true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoCreateTime:false"`
	Password  string    `json:"-" gorm:"size:255;not null"`
}

// Note: Use this func to override the default table-name
func (User) TableName() string {
	return "users"
}

// Note: Create Scope for Query
// Usage: repo.db.Scopes(ActiveUsers).Find(&users)
func ActiveUsers(db *gorm.DB) *gorm.DB {
	return db.Where("active = ?", true)
}
