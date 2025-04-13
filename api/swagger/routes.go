package api

import (
	"go-todos-api/handlers"
	"go-todos-api/middlewares"
	"go-todos-api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(repo *repositories.Repository) *gin.Engine {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Hello World Apis
	helloWorldHandler := new(handlers.HelloWorldHandler)

	// Public routes
	r.GET("/hello-world", helloWorldHandler.SayHelloWorld)
	r.GET("/hello-world-bean/path-variable/:username", helloWorldHandler.SayHelloWorldTo)

	// Authentication
	authHandler := handlers.NewAuthenticationHandler(repo.User)

	// Note: Basic Authentication
	// authoried := r.Group("/", BasicAuthMiddleware)
	// r.GET("/basicauth", authHandler.CheckBasicAuth) // Public route

	// Note: Jwt Authentication
	// Protected routes
	// authoried := r.Group("/")
	// authoried.Use(middlewares.JwtAuthMiddleware()) // All routes in this group require authentication
	authoried := r.Group("/", middlewares.JwtAuthMiddleware()) // Short way to add middleware to group-routes
	// Apply middleware to specific routes that require authentication
	// r.POST("/add", middlewares.JwtAuthMiddleware, func(c *gin.Context) {
	// 	// ...
	// })
	r.POST("/authenticate", authHandler.Login) // Public route
	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	})

	// {
	// 	api.GET("/profile", routes.GetProfile)

	// 	// Note: Use RoleMiddleware for Role-Based Access Control (RBAC)
	// 	// Admin-only routes
	// 	admin := authoried.Group("/admin")
	// 	admin.Use(middlewares.RoleMiddleware("admin")) // Only users with "admin" role can access
	// 	{
	// 		admin.GET("/users", userHandler.GetAllUsers)
	// 		admin.POST("/users", userHandler.CreateUser)
	// 	}
	// }

	// User Apis
	userRoutes := authoried.Group("/")
	{
		userHandler := handlers.NewUserHandler(repo.User)

		r.GET("/users", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.Retrieve)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Todo Apis
	todoRoutes := authoried.Group("/users")
	{
		todoHandler := handlers.NewTodoHandler(repo.Todo)
		todoRoutes.GET("/:username/todos", todoHandler.GetAllUserTodos)
		todoRoutes.POST("/:username/todos", todoHandler.CreateTodo)
		todoRoutes.GET("/:username/todos/:id", todoHandler.Retrieve)
		todoRoutes.PUT("/:username/todos/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:username/todos/:id", todoHandler.DeleteTodo)
	}

	return r
}
