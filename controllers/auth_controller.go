package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wepull/Bug-Tracker/config"
	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/dto/responses"
	"github.com/wepull/Bug-Tracker/services"
	"github.com/wepull/Bug-Tracker/utils"
)

type AuthController struct {
	AuthService *services.AuthService
	Config      *config.Config
}

// POST /auth/register
func (ac *AuthController) Register(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.AuthService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, ac.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	resp := responses.AuthResponse{
		Token: token,
		User: responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	c.JSON(http.StatusCreated, resp)
}

// POST /auth/login
func (ac *AuthController) Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.AuthService.Login(req.UsernameOrEmail, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID, ac.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	resp := responses.AuthResponse{
		Token: token,
		User: responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	c.JSON(http.StatusOK, resp)
}

// POST /auth/logout
// For stateless JWT, you typically just ask the client to discard the token.
func (ac *AuthController) Logout(c *gin.Context) {
	// If you keep a blacklist, mark token as invalid.
	// Otherwise, simply return a success message.
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
