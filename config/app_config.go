package config

import (
	"fmt"
	"go-todos-api/db"
	"go-todos-api/repositories"
	"sync"

	"gorm.io/gorm"
)

// Note: Implement Dependencies Injection
var lock = &sync.Mutex{}
var appConfigInstance *AppConfig

type AppConfig struct {
	DB       *gorm.DB
	UserRepo *repositories.UserRepository
	TodoRepo *repositories.TodoRepository
}

func InitApp() *AppConfig {

	dbInstance := db.InitDb()
	db.SeedDB(dbInstance)

	fmt.Println("Init >>> Seeded DB")

	appConfigInstance = &AppConfig{
		DB:       dbInstance,
		UserRepo: &repositories.UserRepository{DB: dbInstance},
		TodoRepo: &repositories.TodoRepository{DB: dbInstance},
	}
	fmt.Println("Init >>> App")

	return appConfigInstance
}

func GetAppConfig() *AppConfig {
	if appConfigInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if appConfigInstance == nil {
			panic("App configs not initialized")
		}
	}
	return appConfigInstance
}
