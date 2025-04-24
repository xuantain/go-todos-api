package main

import (
	"flag"
	"fmt"
	"os"

	"go-todos-api/config"
	"go-todos-api/dependencies"
	_ "go-todos-api/docs"
	"go-todos-api/router"
)

// @title			Gingo Todos API
// @version			1.0
// @description		A todo management service API in Go using Gin framework.
// @host			localhost
// @BasePath		/api/
func main() {

	// todo: Use .env instead?
	environment := flag.String("env", "dev", "")
	fmt.Println("-env=", *environment)

	flag.Usage = func() {
		fmt.Println("Usage: server -env={mode}")
		os.Exit(1)
	}
	flag.Parse()

	// Get configs from *.yml files
	config.Init(*environment)
	configs := config.GetConfig()

	// Init appConfig
	config.InitApp()

	// Load inner dependencies
	deps := dependencies.Init()

	// Setup routes
	server := router.SetupRoutes(deps, nil)
	server = router.SetupApis(deps, server)

	// Start web-service
	address := configs.GetString("server.address")
	port := configs.GetString("server.port")
	url := address + ":" + port
	server.Run(url)
}
