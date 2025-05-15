package routes

import (
	"go-gin-project/controllers"

	"github.com/gin-gonic/gin"
)

// SetupPingRoutes sets up the ping routes.
func SetupPingRoutes(router *gin.RouterGroup, pingController *controllers.PingController, authMiddleware gin.HandlerFunc) {
	pingRoutes := router.Group("/ping")
	pingRoutes.Use(authMiddleware) // Apply auth middleware to all ping routes
	{
		pingRoutes.GET("", pingController.Ping)
	}
}
