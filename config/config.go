package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config stores all configuration of the application.
// The values are read by godotenv from a .env file or environment variables.
type Config struct {
	AdminPassword string
	JWTSecret     string
	JWTExpiration time.Duration
	DBPath        string
}

var AppConfig *Config

// InitConfig loads configuration from .env file or environment variables.
func InitConfig() {
	// Load .env file if it exists, otherwise use environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin123" // Default admin password if not set
		log.Println("ADMIN_PASSWORD not set, using default: admin123")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key" // Default JWT secret if not set
		log.Println("JWT_SECRET not set, using default: your-secret-key")
	}

	jwtExpHoursStr := os.Getenv("JWT_EXPIRATION_HOURS")
	jwtExpHours := 24 // Default to 24 hours
	if jwtExpHoursStr != "" {
		tempExp, err := time.ParseDuration(jwtExpHoursStr + "h")
		if err == nil {
			jwtExpHours = int(tempExp.Hours())
		} else {
			log.Printf("Invalid JWT_EXPIRATION_HOURS format: %s, using default 24h", jwtExpHoursStr)
		}
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./gogin.db" // Default SQLite path
		log.Println("DB_PATH not set, using default: ./gogin.db")
	}

	AppConfig = &Config{
		AdminPassword: adminPassword,
		JWTSecret:     jwtSecret,
		JWTExpiration: time.Hour * time.Duration(jwtExpHours),
		DBPath:        dbPath,
	}
	log.Println("Configuration loaded successfully")
}
