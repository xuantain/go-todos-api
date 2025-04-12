package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-todos-api/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// Mock data
/*
var todoList = GetTodoListMockData()

func (h TodoRepository) GetAllTodos() []models.Todo {
	return todoList
}

func (h TodoRepository) GetByID(id int) (models.Todo, bool) {
	for _, todo := range todoList {
		if todo.ID == id {
			return todo, true
		}
	}

	return models.Todo{}, false
}

func (h TodoRepository) InsertNewTodo(newTodo models.Todo) bool {
	newTodo.ID = len(todoList) + 1
	todoList = append(todoList, newTodo)

	return true
}

func (h TodoRepository) UpdateTodo(updatedTodo models.Todo) bool {
	idx := slices.IndexFunc(todoList, func(item models.Todo) bool { return item.ID == updatedTodo.ID })
	if idx == -1 {
		return false
	}

	todoList = slices.Replace(todoList, idx, idx+1, updatedTodo)
	return true
}

func (h TodoRepository) DeleteTodo(todoId int) bool {
	todoList = slices.DeleteFunc(todoList, func(item models.Todo) bool { return item.ID == todoId })
	return true
}
*/

// Implement repository methods

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *TodoRepository) FindByID(ctx context.Context, id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.WithContext(ctx).First(&todo, id).Error
	fmt.Println("todo >>>", todo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &todo, err
}

func (r *TodoRepository) Update(ctx context.Context, todo *models.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *TodoRepository) Delete(ctx context.Context, id uint) error {
	// Hard Delete
	// return r.WithContext(ctx).Unscoped().Delete(&models.Todo{}, id).Error
	// Soft Delete
	return r.db.WithContext(ctx).Delete(&models.Todo{}, id).Error
}

func (r *TodoRepository) List(ctx context.Context, page, limit int) ([]*models.Todo, error) {
	var todos []*models.Todo
	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Todo{}).Count(&count).Error
	return count, err
}

func (r *TodoRepository) ListByUsername(ctx context.Context, username string, page, limit int) ([]*models.Todo, error) {
	var todos []*models.Todo
	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		Offset(offset).Limit(limit).
		Find(&todos).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return todos, err
}
