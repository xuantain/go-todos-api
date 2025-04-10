package db

import (
	"go-todos-api/models"
	"log"
	"os"
	"time"

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
		users := []models.User{
			{
				Name:      "Test User 1",
				Username:  "todo",
				Email:     "test1@example.com",
				BirthDay:  "2000-01-01",
				Gender:    "Male",
				PhotoURL:  "",
				LastLogin: time.Now().Add(-2),
				Active:    true,
				UpdatedAt: time.Now().Add(-10),
				Password:  "",
			},
			{
				Name:      "Test User 2",
				Username:  "aha",
				Email:     "test2@example.com",
				BirthDay:  "2001-01-15",
				Gender:    "Male",
				PhotoURL:  "",
				LastLogin: time.Now().Add(-2),
				Active:    true,
				UpdatedAt: time.Now().Add(-10),
				Password:  "",
			},
			{
				Name:      "Test User 13",
				Username:  "oho",
				Email:     "test3@example.com",
				BirthDay:  "2002-03-02",
				Gender:    "Male",
				PhotoURL:  "",
				LastLogin: time.Now().Add(-2),
				Active:    true,
				UpdatedAt: time.Now().Add(-10),
				Password:  "",
			},
		}

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
		todos := []models.Todo{
			{
				UserName:    "todo",
				Description: "Learn AWS",
				TargetDate:  "2025-06-15",
				Done:        false,
			},
			{
				UserName:    "todo",
				Description: "Learn Azure",
				TargetDate:  "2025-07-15",
				Done:        false,
			},
			{
				UserName:    "todo",
				Description: "Learn DevOp",
				TargetDate:  "2025-08-15",
				Done:        false,
			},
			{
				UserName:    "todo",
				Description: "Learn Deno",
				TargetDate:  "2025-09-15",
				Done:        false,
			},
		}

		result := db.Create(&todos)
		if result.Error != nil {
			log.Fatalf("Failed to seed database: %v", result.Error)
		}
		log.Println("Database seeded successfully")
	}
}
