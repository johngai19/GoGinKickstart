package models

import (
	"log"

	"go-gin-project/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the SQLite database connection and runs migrations.
func InitDB() *gorm.DB {
	cfg := config.AppConfig

	// Use the single DBPath from your config
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate your models
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	DB = db
	log.Println("Database initialized and schema migrated.")
	return DB
}

// GetDB returns the global DB instance.
func GetDB() *gorm.DB {
	return DB
}
