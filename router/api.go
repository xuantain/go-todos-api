package router

import (
	"go-todos-api/dependencies"
	"go-todos-api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupApis(deps *dependencies.Dependencies, r *gin.Engine) *gin.Engine {

	if r == nil {
		r = gin.Default()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}

	api := r.Group("/api")

	// Set route for Swagger docs
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Hello World Apis
	helloWorldHandler := deps.HelloHandler

	/** Public routes */

	api.GET("/hello-world", helloWorldHandler.SayHelloWorld)
	api.GET("/hello-world-bean/path-variable/:username", helloWorldHandler.SayHelloWorldTo)

	// Authentication
	authHandler := deps.AuthHandler

	// Note: Basic Authentication
	// authoried := api.Group("/", BasicAuthMiddleware)
	// api.GET("/basicauth", authHandler.CheckBasicAuth) // Public route

	// Note: Jwt Authentication
	// Protected routes
	// authoried := api.Group("/")
	// authoried.Use(middlewares.JwtAuthMiddleware()) // All routes in this group require authentication
	authoried := api.Group("/", middlewares.JwtAuthMiddleware()) // Short way to add middleware to group-routes
	// Apply middleware to specific routes that require authentication
	// api.POST("/add", middlewares.JwtAuthMiddleware, func(c *gin.Context) {
	// 	// ...
	// })
	api.POST("/authenticate", authHandler.Login) // Public route
	api.GET("/logout", func(c *gin.Context) {
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

	/** User APIs */
	userRoutes := authoried.Group("/users")
	{
		userHandler := deps.UserHandler

		api.GET("/users", userHandler.GetAllUsers)
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("/:userId", userHandler.Retrieve)
		userRoutes.PUT("/:userId", userHandler.UpdateUser)
		userRoutes.DELETE("/:userId", userHandler.DeleteUser)
	}

	/** Todo APIs */
	todoRoutes := authoried.Group("/users/:userId")
	{
		todoHandler := deps.TodoHandler
		todoRoutes.GET("/todos", todoHandler.GetAllUserTodos)
		todoRoutes.POST("/todos", todoHandler.CreateTodo)
		todoRoutes.GET("/todos/:todoId", todoHandler.Retrieve)
		todoRoutes.PUT("/todos/:todoId", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/todos/:todoId", todoHandler.DeleteTodo)
	}

	return r
}
