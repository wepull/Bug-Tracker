package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/dto/responses"
	"github.com/wepull/Bug-Tracker/services"
)

type UserController struct {
	UserService *services.UserService
}

// GET /users
func (uc *UserController) ListUsers(c *gin.Context) {
	users, err := uc.UserService.ListUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []responses.UserResponse
	for _, u := range users {
		resp = append(resp, responses.UserResponse{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GET /users/:identifier (could be ID, username, or email)
func (uc *UserController) GetUser(c *gin.Context) {
	identifier := c.Param("identifier")

	// Try parse as ID
	if id, err := strconv.ParseUint(identifier, 10, 64); err == nil {
		user, err := uc.UserService.GetUserByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		resp := responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// If not ID, check if it's an email
	if strings.Contains(identifier, "@") {
		// treat as email
		user, err := uc.UserService.GetUserByEmail(identifier)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		resp := responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// Otherwise treat as username
	user, err := uc.UserService.GetUserByUsername(identifier)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	resp := responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	c.JSON(http.StatusOK, resp)
}

// PUT /users/:id
func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	idVal, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := uc.UserService.UpdateUser(uint(idVal), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.UserResponse{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
	}
	c.JSON(http.StatusOK, resp)
}

// DELETE /users/:id
func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	idVal, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = uc.UserService.DeleteUser(uint(idVal))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
