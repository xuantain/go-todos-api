package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	User UserRepository
	Todo TodoRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: *NewUserRepository(db),
		Todo: *NewTodoRepository(db),
	}
}
