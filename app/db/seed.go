package db

import (
	"go-todos-api/models"
	"go-todos-api/repositories"
	"log"
	"os"

	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	seedUsers(db)
	seedTodos(db)
}

func seedUsers(db *gorm.DB) {
	// AutoMigrate will create tables if they don't exist
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Note: Prevent seeding test data to production
	if os.Getenv("APP_ENV") == "prod" {
		return
	}

	// Check if data already exists
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		// Seed initial data
		users := repositories.GetUserListMockData()

		result := db.Create(&users)
		if result.Error != nil {
			log.Fatalf("Failed to seed database: %v", result.Error)
		}
		log.Println("Database seeded successfully")
	}
}

func seedTodos(db *gorm.DB) {
	// AutoMigrate will create tables if they don't exist
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Note: Prevent seeding test data to production
	if os.Getenv("APP_ENV") == "prod" {
		return
	}

	// Check if data already exists
	var count int64
	db.Model(&models.Todo{}).Count(&count)

	if count == 0 {
		// Seed initial data
		todos := repositories.GetTodoListMockData()

		result := db.Create(&todos)
		if result.Error != nil {
			log.Fatalf("Failed to seed database: %v", result.Error)
		}
		log.Println("Database seeded successfully")
	}
}
