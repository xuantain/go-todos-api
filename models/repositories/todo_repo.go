package models

import (
	"go-todos-api/models"
	"slices"
)

type TodoRepo struct{}

var todoList = GetTodoListMockData()

func (h TodoRepo) GetAllTodos() []models.Todo {
	return todoList
}

func (h TodoRepo) GetByID(id int) (models.Todo, bool) {
	for _, todo := range todoList {
		if todo.ID == id {
			return todo, true
		}
	}

	return models.Todo{}, false
}

func (h TodoRepo) InsertNewTodo(newTodo models.Todo) bool {
	newTodo.ID = len(todoList) + 1
	todoList = append(todoList, newTodo)

	return true
}

func (h TodoRepo) UpdateTodo(updatedTodo models.Todo) bool {
	idx := slices.IndexFunc(todoList, func(item models.Todo) bool { return item.ID == updatedTodo.ID })
	if idx == -1 {
		return false
	}

	todoList = slices.Replace(todoList, idx, idx+1, updatedTodo)
	return true
}

func (h TodoRepo) DeleteTodo(todoId int) bool {
	todoList = slices.DeleteFunc(todoList, func(item models.Todo) bool { return item.ID == todoId })
	return true
}
