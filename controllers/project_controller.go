package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/dto/responses"
	"github.com/wepull/Bug-Tracker/services"
)

type ProjectController struct {
	ProjectService *services.ProjectService
}

// POST /projects
func (pc *ProjectController) CreateProject(c *gin.Context) {
	var req requests.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := pc.ProjectService.CreateProject(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserID:      project.UserID,
		TeamID:      project.TeamID,
	}
	c.JSON(http.StatusCreated, resp)
}

// GET /projects/:projectId
func (pc *ProjectController) GetProject(c *gin.Context) {
	pidParam := c.Param("projectId")
	pidVal, err := strconv.ParseUint(pidParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	project, err := pc.ProjectService.GetProjectByID(uint(pidVal))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := responses.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserID:      project.UserID,
		TeamID:      project.TeamID,
	}
	c.JSON(http.StatusOK, resp)
}

// PUT /projects/:projectId
func (pc *ProjectController) UpdateProject(c *gin.Context) {
	pidParam := c.Param("projectId")
	pidVal, err := strconv.ParseUint(pidParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var req requests.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProject, err := pc.ProjectService.UpdateProject(uint(pidVal), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.ProjectResponse{
		ID:          updatedProject.ID,
		Name:        updatedProject.Name,
		Description: updatedProject.Description,
		UserID:      updatedProject.UserID,
		TeamID:      updatedProject.TeamID,
	}
	c.JSON(http.StatusOK, resp)
}

// DELETE /projects/:projectId
func (pc *ProjectController) DeleteProject(c *gin.Context) {
	pidParam := c.Param("projectId")
	pidVal, err := strconv.ParseUint(pidParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	err = pc.ProjectService.DeleteProject(uint(pidVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}

// GET /users/:userId/projects
func (pc *ProjectController) ListUserProjects(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDVal, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	projects, err := pc.ProjectService.ListUserProjects(uint(userIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := []responses.ProjectResponse{}
	for _, p := range projects {
		resp = append(resp, responses.ProjectResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			UserID:      p.UserID,
			TeamID:      p.TeamID,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GET /teams/:teamId/projects
func (pc *ProjectController) ListTeamProjects(c *gin.Context) {
	teamIDParam := c.Param("teamId")
	teamIDVal, err := strconv.ParseUint(teamIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team id"})
		return
	}

	projects, err := pc.ProjectService.ListTeamProjects(uint(teamIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := []responses.ProjectResponse{}
	for _, p := range projects {
		resp = append(resp, responses.ProjectResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			UserID:      p.UserID,
			TeamID:      p.TeamID,
		})
	}
	c.JSON(http.StatusOK, resp)
}
