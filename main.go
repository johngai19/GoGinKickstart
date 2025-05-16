package main

import (
	"log"

	"go-gin-project/config"
	"go-gin-project/controllers"
	"go-gin-project/middlewares"
	"go-gin-project/models"
	"go-gin-project/routes"
	"go-gin-project/services"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
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
	defer db.Close()

	// Create initial admin user if not exists
	userService := services.NewUserService(db)
	err = userService.CreateInitialAdminUser()
	if err != nil {
		log.Fatalf("Failed to create initial admin user: %v", err)
	}

	r := gin.Default()
	//Todo Allow CORS for all origins, methods, and headers, Must be changed in production


	r.Use(cors.Default())


	// Serve static files (HTML, CSS, JS for login and ping pong test)
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/login", func(c *gin.Context) {
		c.File("./static/login.html")
	})

	apiV1 := r.Group("/api/v1")
	{
		authController := controllers.NewAuthController(db)
		pingController := controllers.NewPingController()

		routes.SetupAuthRoutes(apiV1, authController)
		routes.SetupPingRoutes(apiV1, pingController, middlewares.AuthMiddleware())
		routes.SetupSwaggerRoutes(r) // Swagger routes on root path
	}

	
	log.Printf("Server starting on port %s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

