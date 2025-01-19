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

type BugController struct {
	BugService *services.BugService
}

// POST /projects/:projectId/bugs
func (bc *BugController) CreateBug(c *gin.Context) {
	projectIDParam := c.Param("projectId")
	projectIDVal, err := strconv.ParseUint(projectIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var req requests.CreateBugRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := middlewares.GetUserIDFromContext(c)

	bug, err := bc.BugService.CreateBug(uint(projectIDVal), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.BugResponse{
		ID:          bug.ID,
		Title:       bug.Title,
		Description: bug.Description,
		Severity:    bug.Severity,
		Status:      bug.Status,
		ProjectID:   bug.ProjectID,
		CreatedBy:   bug.CreatedBy,
		AssignedTo:  bug.AssignedTo,
	}
	c.JSON(http.StatusCreated, resp)
}

// GET /bugs/:bugId
func (bc *BugController) GetBug(c *gin.Context) {
	bugIDParam := c.Param("bugId")
	bugIDVal, err := strconv.ParseUint(bugIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bug id"})
		return
	}

	bug, err := bc.BugService.GetBugByID(uint(bugIDVal))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := responses.BugResponse{
		ID:          bug.ID,
		Title:       bug.Title,
		Description: bug.Description,
		Severity:    bug.Severity,
		Status:      bug.Status,
		ProjectID:   bug.ProjectID,
		CreatedBy:   bug.CreatedBy,
		AssignedTo:  bug.AssignedTo,
	}
	c.JSON(http.StatusOK, resp)
}

// PUT /bugs/:bugId
func (bc *BugController) UpdateBug(c *gin.Context) {
	bugIDParam := c.Param("bugId")
	bugIDVal, err := strconv.ParseUint(bugIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bug id"})
		return
	}

	var req requests.UpdateBugRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBug, err := bc.BugService.UpdateBug(uint(bugIDVal), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := responses.BugResponse{
		ID:          updatedBug.ID,
		Title:       updatedBug.Title,
		Description: updatedBug.Description,
		Severity:    updatedBug.Severity,
		Status:      updatedBug.Status,
		ProjectID:   updatedBug.ProjectID,
		CreatedBy:   updatedBug.CreatedBy,
		AssignedTo:  updatedBug.AssignedTo,
	}
	c.JSON(http.StatusOK, resp)
}

// DELETE /bugs/:bugId
func (bc *BugController) DeleteBug(c *gin.Context) {
	bugIDParam := c.Param("bugId")
	bugIDVal, err := strconv.ParseUint(bugIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bug id"})
		return
	}

	if err := bc.BugService.DeleteBug(uint(bugIDVal)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bug deleted"})
}

// GET /bugs/assigned/:userId
func (bc *BugController) ListBugsAssigned(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDVal, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	bugs, err := bc.BugService.ListBugsAssignedTo(uint(userIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []responses.BugResponse
	for _, b := range bugs {
		resp = append(resp, responses.BugResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Severity:    b.Severity,
			Status:      b.Status,
			ProjectID:   b.ProjectID,
			CreatedBy:   b.CreatedBy,
			AssignedTo:  b.AssignedTo,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GET /bugs/created/:userId
func (bc *BugController) ListBugsCreated(c *gin.Context) {
	userIDParam := c.Param("userId")
	userIDVal, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	bugs, err := bc.BugService.ListBugsCreatedBy(uint(userIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []responses.BugResponse
	for _, b := range bugs {
		resp = append(resp, responses.BugResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Severity:    b.Severity,
			Status:      b.Status,
			ProjectID:   b.ProjectID,
			CreatedBy:   b.CreatedBy,
			AssignedTo:  b.AssignedTo,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// GET /projects/:projectId/bugs
func (bc *BugController) ListBugsInProject(c *gin.Context) {
	projectIDParam := c.Param("projectId")
	projectIDVal, err := strconv.ParseUint(projectIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	bugs, err := bc.BugService.ListBugsInProject(uint(projectIDVal))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []responses.BugResponse
	for _, b := range bugs {
		resp = append(resp, responses.BugResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Severity:    b.Severity,
			Status:      b.Status,
			ProjectID:   b.ProjectID,
			CreatedBy:   b.CreatedBy,
			AssignedTo:  b.AssignedTo,
		})
	}
	c.JSON(http.StatusOK, resp)
}
