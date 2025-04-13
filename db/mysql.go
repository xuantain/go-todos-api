package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	DB_USERNAME, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT string
}

var Db *gorm.DB

func InitDb() *gorm.DB {
	fmt.Println("InitApp >>")

	Db = connectSqlDB()

	// err := Db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	return nil
	// }

	return Db
}

func connectSqlDB() *gorm.DB {
	fmt.Println("InitApp >> connectSqlDB()")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		return nil
	}

	return db
}

func isConnected() bool {
	// res, err := Db.Call("1+1")
	// if err != nil {
	// 	return false
	// }
	// return res.Data == int64(2)
	return true
}

func Conn() *gorm.DB {

	if Db == nil {
		fmt.Print("Connecting to DB...")
		Db = connectSqlDB()
	}

	connected := isConnected()
	for !connected {
		fmt.Print("Connection to DB was lost. Waiting 5s...")
		time.Sleep(5 * time.Second)
		fmt.Print("Reconnecting...")
		Db = connectSqlDB()
		connected = isConnected()
	}

	return Db
}
