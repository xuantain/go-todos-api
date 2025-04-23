package controllers

import (
	"net/http"

	_ "go-todos-api/docs"
	"go-todos-api/models"
	"go-todos-api/repositories"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func (u TodoController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "Index", gin.H{
		"title": "Home Page",
	})
	c.Abort()
}

func (u TodoController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "Welcome", gin.H{
		"welcome":   "Welcome to Todo API docs!",
		"linkTitle": "Click here to access API docs",
		"link":      "/api/docs/index.html",
	})
	c.Abort()
}

func (u TodoController) ListTodos(c *gin.Context) {
	todoList := repositories.GetTodoListMockData()
	c.HTML(http.StatusOK, "TodoList", gin.H{
		"title": "Todos Page",
		"todos": todoList,
	})
	c.Abort()
}

func (u TodoController) CreateTodo(c *gin.Context) {
	newTodo := models.Todo{}
	c.HTML(http.StatusOK, "Todo", gin.H{
		"title": "Todo",
		"todo":  newTodo,
	})
	c.Abort()
}

func (u TodoController) UpdateTodo(c *gin.Context) {
	todo := repositories.GetTodoListMockData()[0]
	c.HTML(http.StatusOK, "Todo", gin.H{
		"title": "Todo",
		"todo":  todo,
	})
	c.Abort()
}
