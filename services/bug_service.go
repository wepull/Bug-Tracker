package services

import (
	"errors"

	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/models"
	"gorm.io/gorm"
)

type BugService struct {
	DB *gorm.DB
}

func (s *BugService) CreateBug(projectID, creatorID uint, req requests.CreateBugRequest) (*models.Bug, error) {
	bug := &models.Bug{
		Title:       req.Title,
		Description: req.Description,
		Severity:    req.Severity,
		Status:      "open",
		ProjectID:   projectID,
		CreatedBy:   creatorID,
	}
	if req.AssignedTo != nil {
		bug.AssignedTo = *req.AssignedTo
	}
	if err := s.DB.Create(bug).Error; err != nil {
		return nil, err
	}
	return bug, nil
}

func (s *BugService) GetBugByID(bugID uint) (*models.Bug, error) {
	var bug models.Bug
	if err := s.DB.First(&bug, bugID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("bug not found")
		}
		return nil, err
	}
	return &bug, nil
}

func (s *BugService) UpdateBug(bugID uint, req requests.UpdateBugRequest) (*models.Bug, error) {
	bug, err := s.GetBugByID(bugID)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		bug.Title = *req.Title
	}
	if req.Description != nil {
		bug.Description = *req.Description
	}
	if req.Severity != nil {
		bug.Severity = *req.Severity
	}
	if req.Status != nil {
		bug.Status = *req.Status
	}
	if req.AssignedTo != nil {
		bug.AssignedTo = *req.AssignedTo
	}

	if err := s.DB.Save(bug).Error; err != nil {
		return nil, err
	}
	return bug, nil
}

func (s *BugService) DeleteBug(bugID uint) error {
	bug, err := s.GetBugByID(bugID)
	if err != nil {
		return err
	}
	return s.DB.Delete(bug).Error
}

func (s *BugService) ListBugsAssignedTo(userID uint) ([]models.Bug, error) {
	var bugs []models.Bug
	err := s.DB.Where("assigned_to = ?", userID).Find(&bugs).Error
	return bugs, err
}

func (s *BugService) ListBugsCreatedBy(userID uint) ([]models.Bug, error) {
	var bugs []models.Bug
	err := s.DB.Where("created_by = ?", userID).Find(&bugs).Error
	return bugs, err
}

func (s *BugService) ListBugsInProject(projectID uint) ([]models.Bug, error) {
	var bugs []models.Bug
	err := s.DB.Where("project_id = ?", projectID).Find(&bugs).Error
	return bugs, err
}
