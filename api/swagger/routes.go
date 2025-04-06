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

	r.GET("basicauth", authHandler.GetBasicAuth)
	r.GET("authenticate", authHandler.Authenticate)

	// User Apis
	userHandler := new(handlers.UserHandler)

	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.Retrieve)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// Todo Apis
	todoHandler := new(handlers.TodoHandler)

	r.GET("/todos", todoHandler.GetAllTodos)
	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos/:id", todoHandler.Retrieve)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
