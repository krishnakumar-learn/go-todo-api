package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Connect establishes a connection to the MySQL database using GORM.
// It returns a pointer to the gorm.DB instance.
// If the connection fails, the application will log the error and exit.
func Connect() *gorm.DB {
	dsn := "todo:todopass@tcp(192.168.1.49:3306)/go_todo_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
