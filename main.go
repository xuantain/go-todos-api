package main

import (
	"flag"
	"fmt"
	"os"

	router "go-todos-api/api/swagger"
	"go-todos-api/config"
	"go-todos-api/dependencies"
)

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
	server := router.SetupRoutes(deps)

	// Start web-service
	address := configs.GetString("server.address")
	port := configs.GetString("server.port")
	url := address + ":" + port
	server.Run(url)
}
