package models

import (
	"go-todos-api/pkg/types"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `json:"id,string"      form:"id,string"      gorm:"primary_key"`
	UserID      uint           `json:"userId,string"  form:"userId,string"  gorm:"foreign_key;not null"`
	Description string         `json:"description"    form:"description"    gorm:"size:255;not null"     binding:"required"`
	TargetDate  types.DateType `json:"targetDate"     form:"targetDate"     gorm:"type:date"             binding:"required"`
	Done        bool           `json:"done"           form:"done"           gorm:"default:false"`
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
