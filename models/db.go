package models

import (
	"log"

	"go-gin-project/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// InitDB initializes the database connection.
func InitDB() *gorm.DB {
	dbPath := config.AppConfig.DBPath
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		// Ensure the directory exists if using a relative path like ./data/gogin.db
		// For simplicity, this example uses a root-level db file.
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
