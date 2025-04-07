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
	userRoutes := authoried.Group("/")
	{
		userHandler := new(handlers.UserHandler)

		r.GET("/users", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.Retrieve)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Todo Apis
	todoRoutes := authoried.Group("/users")
	{
		todoHandler := new(handlers.TodoHandler)

		todoRoutes.GET("/:username/todos", todoHandler.GetAllTodos)
		todoRoutes.POST("/:username/todos", todoHandler.CreateTodo)
		todoRoutes.GET("/:username/todos/:id", todoHandler.Retrieve)
		todoRoutes.PUT("/:username/todos/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:username/todos/:id", todoHandler.DeleteTodo)
	}

}
