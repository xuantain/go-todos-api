package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint      `json:"id,string"   form:"id,string"   gorm:"primary_key"`
	Username    string    `json:"username"    form:"username"    gorm:"size:30;not null"`
	Description string    `json:"description" form:"description" gorm:"size:255;not null"    binding:"required"`
	TargetDate  time.Time `json:"targetDate"  form:"targetDate"  gorm:"autoCreateTime:false" binding:"required"`
	Done        bool      `json:"done"        form:"done"        gorm:"default:true"`
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
