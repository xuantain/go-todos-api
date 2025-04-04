package models

import "go-todos-api/models"

type TodoRepo struct{}

func (h TodoRepo) GetAllTodos() []models.Todo {
	return GetTodoListMockData()
}

func (h TodoRepo) GetByID(id int) *models.Todo {
	// var todo *models.Todo
	todoList := GetTodoListMockData()

	for _, user := range todoList {
		if user.ID == id {
			return &user
		}
	}

	return nil
}
