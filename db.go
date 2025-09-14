package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&User{}, &Book{})
	return db
}
