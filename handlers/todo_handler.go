package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-todos-api/models"
	repos "go-todos-api/models/repositories"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct{}

var todoRepo = new(repos.TodoRepo)

func (todo TodoHandler) GetAllTodos(c *gin.Context) {
	if username := c.Param("username"); username != "" {
		todoList := todoRepo.GetAllTodos()

		userTodoList := []models.Todo{}

		for _, item := range todoList {
			if item.UserName == username {
				userTodoList = append(userTodoList, item)
			}
		}

		c.JSON(http.StatusOK, userTodoList)
		return
	}
	c.Abort()
}

func (todo TodoHandler) CreateTodo(c *gin.Context) {
	newTodo := models.Todo{}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !todoRepo.InsertNewTodo(newTodo) {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"todo": newTodo})
}

func (u TodoHandler) UpdateTodo(c *gin.Context) {
	if c.Param("id") != "" {

		todoId, _ := strconv.Atoi(c.Param("id"))
		updatedTodo := models.Todo{}

		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "todo": updatedTodo})
			return
		}

		updatedTodo.ID = todoId
		if !todoRepo.UpdateTodo(updatedTodo) {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Cannot update"})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"todo": updatedTodo})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u TodoHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		todoId, _ := strconv.Atoi(c.Param("id"))
		todo, isExist := todoRepo.GetByID(todoId)

		if !isExist {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo {%d} does not exist", todoId)})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, todo)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u TodoHandler) DeleteTodo(c *gin.Context) {
	if c.Param("id") != "" {

		todoId, _ := strconv.Atoi(c.Param("id"))
		isGone := todoRepo.DeleteTodo(todoId)

		if !isGone {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot delete Todo {%d}", todoId)})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Deleted todo id:" + c.Param("id")})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
