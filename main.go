package main

import (
	"log"

	"go-gin-project/config"
	"go-gin-project/models"
	"go-gin-project/routes"
	"go-gin-project/services"

	"github.com/joho/godotenv"
)

// @title Go Gin Project API
// @version 1.0
// @description This is a sample server for a Go Gin project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default or environment variables")
	}

	config.InitConfig()
	db := models.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create initial admin user if not exists
	userService := services.NewUserService(db)
	if err := userService.CreateInitialAdminUser(); err != nil {
		log.Fatalf("Failed to create initial admin user: %v", err)
	}

	r := routes.NewRouter(db)

	log.Printf("Server starting on port %s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

