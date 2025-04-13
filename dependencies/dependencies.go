package dependencies

import (
	"fmt"
	"go-todos-api/config"
	"go-todos-api/handlers"
)

type Dependencies struct {
	HelloHandler *handlers.HelloWorldHandler
	AuthHandler  *handlers.AuthenticationHandler
	UserHandler  *handlers.UserHandler
	TodoHandler  *handlers.TodoHandler
}

func Init() *Dependencies {
	fmt.Println("Dependencies >>> Init() >>")
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
	}
}
