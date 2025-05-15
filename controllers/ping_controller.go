package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingController handles ping requests.
type PingController struct{}

// NewPingController creates a new PingController.
func NewPingController() *PingController {
	return &PingController{}
}

// PingSuccessResponse defines the structure for a successful ping response.
// swagger:model PingSuccessResponse
type PingSuccessResponse struct {
	Message string `json:"message" example:"pong"`
	User    string `json:"user,omitempty" example:"admin"` // User who made the request
}

// Ping godoc
// @Summary Ping endpoint
// @Description Returns a pong message, requires authentication.
// @Tags ping
// @Produce  json
// @Success 200 {object} PingSuccessResponse
// @Failure 401 {object} GenericErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /ping [get]
func (pc *PingController) Ping(c *gin.Context) {
	username, exists := c.Get("username")
	userStr := ""
	if exists {
		userStr, _ = username.(string)
	}
	c.JSON(http.StatusOK, PingSuccessResponse{Message: "pong", User: userStr})
}

