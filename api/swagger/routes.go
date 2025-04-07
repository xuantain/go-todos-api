package api

import (
	"go-todos-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Hello World Apis
	helloWorldHandler := new(handlers.HelloWorldHandler)

	r.GET("/hello-word", helloWorldHandler.SayHelloWorld)
	r.GET("/hello-word/:username", helloWorldHandler.SayHelloWorldTo)

	// Authentication
	authHandler := new(handlers.AuthenticationHandler)

	// Note: Create a route-group with specific middlewares
	r.GET("/authenticate", authHandler.Authenticate)
	authoried := r.Group("/", gin.BasicAuth(gin.Accounts{
		"todo": "aaa",
	}))

	authoried.GET("/basicauth", authHandler.CheckBasicAuth)

	// User Apis
	userHandler := new(handlers.UserHandler)

	r.GET("/users", userHandler.GetAllUsers)
	authoried.POST("/users", userHandler.CreateUser)
	authoried.GET("/users/:id", userHandler.Retrieve)
	authoried.PUT("/users/:id", userHandler.UpdateUser)
	authoried.DELETE("/users/:id", userHandler.DeleteUser)

	// Todo Apis
	todoHandler := new(handlers.TodoHandler)

	authoried.GET("/todos", todoHandler.GetAllTodos)
	authoried.POST("/todos", todoHandler.CreateTodo)
	authoried.GET("/todos/:id", todoHandler.Retrieve)
	authoried.PUT("/todos/:id", todoHandler.UpdateTodo)
	authoried.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
