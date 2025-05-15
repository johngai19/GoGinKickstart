package routes

import (
	"go-gin-project/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes sets up the authentication routes.
func SetupAuthRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
}

