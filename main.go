package main

import (
	"flag"
	"fmt"
	"os"

	api "go-todos-api/api/swagger"
	"go-todos-api/config"
	"go-todos-api/db"

	"github.com/gin-gonic/gin"
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

	config.Init(*environment)
	config := config.GetConfig()

	server := gin.Default()
	api.SetupRoutes(server)

	database := db.InitDb()
	db.SeedDB(database)

	// server.Use(middlewares.AuthMiddleware())
	// server.Use(middlewares.JwtAuthMiddleware())
	// server.Use(middlewares.CORSMiddleware())

	address := config.GetString("server.address")
	port := config.GetString("server.port")

	url := address + ":" + port

	server.Run(url)
}
