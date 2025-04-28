package router

import (
	"go-todos-api/dependencies"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(deps *dependencies.Dependencies, r *gin.Engine) *gin.Engine {

	if r == nil {
		r = gin.Default()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}

	// Note: The root-path is the project-folder
	//   We can use root-path => "./" or "" (empty)
	// Load all html/tmpl templates from path
	// r.LoadHTMLGlob("templates/**/*") // Work as well
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static/", "./assets/static")

	todoController := deps.TodoController

	r.GET("/", todoController.Index)
	r.GET("/welcome", todoController.Welcome)
	r.GET("/todos", todoController.ListTodos)
	r.GET("/todo", todoController.CreateTodo)
	r.GET("/todo/:todoId", todoController.UpdateTodo)

	return r
}
