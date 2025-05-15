package controllers

import (
	"net/http"
	"time"

	"go-gin-project/config"
	"go-gin-project/middlewares" // Added missing import for models.User
	"go-gin-project/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

// AuthController handles authentication-related requests.
type AuthController struct {
	userService *services.UserService
}

// NewAuthController creates a new AuthController.
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{userService: services.NewUserService(db)}
}

// LoginPayload defines the expected request body for login.
// swagger:model LoginPayload
type LoginPayload struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"admin123"`
}

// LoginSuccessResponse defines the structure for a successful login response.
// swagger:model LoginSuccessResponse
type LoginSuccessResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// GenericErrorResponse defines the structure for a generic error response.
// swagger:model GenericErrorResponse
type GenericErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"`
}

// Login godoc
// @Summary Log in a user
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   credentials  body LoginPayload true "Login Credentials"
// @Success 200 {object} LoginSuccessResponse
// @Failure 400 {object} GenericErrorResponse "Invalid input"
// @Failure 401 {object} GenericErrorResponse "Invalid credentials"
// @Failure 500 {object} GenericErrorResponse "Internal server error"
// @Router /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Error: "Invalid input: " + err.Error()})
		return
	}

	user, err := ac.userService.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GenericErrorResponse{Error: "Invalid credentials"})
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, GenericErrorResponse{Error: "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(config.AppConfig.JWTExpiration)
	claims := &middlewares.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Error: "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{Token: tokenString})
}
