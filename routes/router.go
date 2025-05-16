package routes

import (
    "go-gin-project/controllers"
    "go-gin-project/middlewares"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
    // 1) switch to release mode (removes the debug‐mode warning)
    gin.SetMode(gin.ReleaseMode)

    // 2) use New() and attach only one Logger+Recovery
    r := gin.New()
    r.Use(
        gin.Logger(),
        gin.Recovery(),
        cors.Default(),
    )

    // --- static / HTML ---
    r.Static("/static", "./static")
    r.GET("/", func(c *gin.Context) { c.File("./static/index.html") })
    r.GET("/login", func(c *gin.Context) { c.File("./static/login.html") })

    // --- API v1 ---
    api := r.Group("/api/v1")
    {
        authController := controllers.NewAuthController(db)
        pingController := controllers.NewPingController()

        SetupAuthRoutes(api, authController)
        SetupPingRoutes(api, pingController, middlewares.AuthMiddleware())
    }

    // --- Swagger UI ---
    SetupSwaggerRoutes(r)

    return r
}