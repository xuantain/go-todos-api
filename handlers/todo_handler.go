package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-todos-api/models"
	"go-todos-api/repositories"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoRepo repositories.TodoRepository
}

// var todoRepo = new(repositories.TodoRepository)

func (h TodoHandler) GetAllUserTodos(c *gin.Context) {

	if username := c.Param("username"); username != "" {

		todoList, err := h.TodoRepo.ListByUsername(c.Request.Context(), username, 0, 10)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todoList)
		return
	}
	c.Abort()
}

func (h TodoHandler) CreateTodo(c *gin.Context) {
	newTodo := models.Todo{}

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.TodoRepo.Create(c.Request.Context(), &newTodo); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"todo": newTodo})
}

func (h TodoHandler) UpdateTodo(c *gin.Context) {

	if c.Param("id") != "" {

		u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		}
		todoId := uint(u64)
		updatedTodo := models.Todo{}

		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "todo": updatedTodo})
			return
		}

		updatedTodo.ID = todoId
		if err := h.TodoRepo.Update(c.Request.Context(), &updatedTodo); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Cannot update"})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"todo": updatedTodo})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (h TodoHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		}
		todoId := uint(u64)
		todo, err := h.TodoRepo.FindByID(c.Request.Context(), todoId)

		if err != nil {
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

func (h TodoHandler) DeleteTodo(c *gin.Context) {

	if c.Param("id") != "" {

		u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		}
		todoId := uint(u64)

		if err := h.TodoRepo.Delete(c.Request.Context(), todoId); err != nil {
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
