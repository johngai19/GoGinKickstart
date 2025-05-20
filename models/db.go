package models

import (
	"log"

	"go-gin-project/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection.
func InitDB() *gorm.DB {
	dbPath := config.AppConfig.DBPath
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	DB = db
	log.Println("Database initialized and schema migrated.")
	return DB
}

// GetDB returns the current database instance.
func GetDB() *gorm.DB {
	return DB
}
