package dependencies

import (
	"fmt"
	"go-todos-api/config"
	"go-todos-api/controllers"
	"go-todos-api/handlers"
)

type Dependencies struct {
	HelloHandler   *handlers.HelloWorldHandler
	AuthHandler    *handlers.AuthenticationHandler
	UserHandler    *handlers.UserHandler
	TodoHandler    *handlers.TodoHandler
	TodoController *controllers.TodoController
}

func Init() *Dependencies {
	fmt.Println("Init >>> Dependencies")
	appConfig := config.GetAppConfig()

	return &Dependencies{
		HelloHandler: &handlers.HelloWorldHandler{},
		AuthHandler: &handlers.AuthenticationHandler{
			UserRepo: *appConfig.UserRepo,
		},
		UserHandler: &handlers.UserHandler{
			UserRepo: *appConfig.UserRepo,
		},
		TodoHandler: &handlers.TodoHandler{
			TodoRepo: *appConfig.TodoRepo,
		},
		TodoController: &controllers.TodoController{},
	}
}
