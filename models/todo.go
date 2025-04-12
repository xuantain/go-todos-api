package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Username    string    `json:"username" gorm:"size:30;not null"`
	Description string    `json:"description" gorm:"size:255;not null"`
	TargetDate  time.Time `json:"targetDate" gorm:"autoCreateTime:false"`
	Done        bool      `json:"done" gorm:"default:true"`
}

// Note: Use this func to override the default table-name
func (Todo) TableName() string {
	return "todos"
}

// Note: Create Scope for Query
// Usage: repo.db.Scopes(TodosDone).Find(&users)
func TodosDone(db *gorm.DB) *gorm.DB {
	return db.Where("done = ?", true)
}
