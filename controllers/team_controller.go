package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/dto/responses"
	"github.com/wepull/Bug-Tracker/middlewares"
	"github.com/wepull/Bug-Tracker/services"
)

type TeamController struct {
	TeamService *services.TeamService
}

// POST /teams
func (tc *TeamController) CreateTeam(c *gin.Context) {
	var req requests.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := middlewares.GetUserIDFromContext(c)

	team, err := tc.TeamService.CreateTeam(req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		CreatedBy:   team.CreatedBy,
	}
	c.JSON(http.StatusCreated, resp)
}

// GET /users/:userId/teams
func (tc *TeamController) ListUserTeams(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDVal, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	teams, err := tc.TeamService.ListUserTeams(uint(userIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []responses.TeamResponse
	for _, t := range teams {
		resp = append(resp, responses.TeamResponse{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
			CreatedBy:   t.CreatedBy,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// PUT /teams/:teamId
func (tc *TeamController) UpdateTeam(c *gin.Context) {
	teamIDParam := c.Param("teamId")
	teamIDVal, err := strconv.ParseUint(teamIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	var req requests.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTeam, err := tc.TeamService.UpdateTeam(uint(teamIDVal), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.TeamResponse{
		ID:          updatedTeam.ID,
		Name:        updatedTeam.Name,
		Description: updatedTeam.Description,
		CreatedBy:   updatedTeam.CreatedBy,
	}
	c.JSON(http.StatusOK, resp)
}

// DELETE /teams/:teamId
func (tc *TeamController) DeleteTeam(c *gin.Context) {
	teamIDParam := c.Param("teamId")
	teamIDVal, err := strconv.ParseUint(teamIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	if err := tc.TeamService.DeleteTeam(uint(teamIDVal)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}

// DELETE /teams/:teamId/members/:userId (leave team)
func (tc *TeamController) RemoveUserFromTeam(c *gin.Context) {
	teamIDParam := c.Param("teamId")
	teamIDVal, err := strconv.ParseUint(teamIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	userIDParam := c.Param("userId")
	userIDVal, err2 := strconv.ParseUint(userIDParam, 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := tc.TeamService.RemoveUserFromTeam(uint(teamIDVal), uint(userIDVal)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed from team"})
}
