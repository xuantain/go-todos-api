package router

import (
	"go-todos-api/dependencies"
	"go-todos-api/middlewares"

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
	authController := deps.AuthController

	web := r.Group("/")
	{
		// Routes requiring authentication
		auth := web.Group("/")
		auth.Use(middlewares.JwtSSAuthMiddleware())
		{
			auth.GET("/welcome", todoController.Welcome)
			auth.GET("/todos", todoController.ListTodos)
			// auth.POST("/todos", todoController.DeleteTodo)
			auth.GET("/todo", todoController.CreateTodo)
			auth.POST("/todo", todoController.CreateTodo)
			auth.GET("/todo/:todoId", todoController.UpdateTodo)
			auth.POST("/todo/:todoId", todoController.UpdateTodo)
			auth.GET("/logout", authController.Logout)
		}

		// Routes for guests (not logged in)
		guest := web.Group("/")
		guest.Use(middlewares.Guest)
		{
			guest.GET("/", todoController.Index)
			guest.GET("/login", todoController.Index)
			guest.POST("/login", authController.Login)
			// guest.GET("/register", todoController.Register)
			// guest.POST("/register", todoController.StoreUser)
		}
	}

	return r
}
