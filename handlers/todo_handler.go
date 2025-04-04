package handlers

import (
	"net/http"

	"go-todos-api/models"
	repos "go-todos-api/models/repositories"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct{}

var todoRepo = new(repos.TodoRepo)

func (todo TodoHandler) GetAllTodos(c *gin.Context) {
	todoList := todoRepo.GetAllTodos()

	c.JSON(http.StatusOK, gin.H{"data": todoList})
	c.Abort()
}

func (todo TodoHandler) CreateTodo(c *gin.Context) {
	// newTodo := c.Param("")
	newTodo := models.Todo{}

	c.JSON(http.StatusCreated, gin.H{"data": newTodo})
	c.Abort()
}

func (u TodoHandler) UpdateTodo(c *gin.Context) {

	newTodo := models.User{}

	c.JSON(http.StatusAccepted, gin.H{"data": newTodo})
	c.Abort()
}

func (u TodoHandler) Retrieve(c *gin.Context) {

	if c.Param("id") != "" {

		todo := todoRepo.GetByID(c.GetInt("id"))

		if todo != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve todo", "error": nil})
			c.Abort()
			return
		}

		c.JSON(http.StatusFound, gin.H{"message": "Todo founded!", "todo": todo})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort() // to stop current handler and ignore calling remaining handlers
}

func (u TodoHandler) DeleteTodo(c *gin.Context) {

	if c.Param("id") != "" {
		c.JSON(http.StatusGone, gin.H{"message": "Deleted todo id:" + c.Param("id")})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}
