package api

import (
	"go-todos-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	helloWorldHandler := new(handlers.HelloWorldHandler)

	r.GET("/hello-word", helloWorldHandler.SayHelloWorld)
	r.GET("/hello-word/:username", helloWorldHandler.SayHelloWorldTo)

	userHandler := new(handlers.UserHandler)

	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.Retrieve)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
}
