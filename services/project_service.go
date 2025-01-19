package services

import (
	"errors"

	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	DB *gorm.DB
}

func (s *ProjectService) CreateProject(req requests.CreateProjectRequest) (*models.Project, error) {
	project := &models.Project{
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
		TeamID:      req.TeamID,
	}
	if err := s.DB.Create(&project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) GetProjectByID(projectID uint) (*models.Project, error) {
	var project models.Project
	if err := s.DB.First(&project, projectID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return &project, nil
}

func (s *ProjectService) UpdateProject(projectID uint, req requests.UpdateProjectRequest) (*models.Project, error) {
	project, err := s.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		project.Name = *req.Name
	}
	if req.Description != nil {
		project.Description = *req.Description
	}
	if err := s.DB.Save(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) DeleteProject(projectID uint) error {
	project, err := s.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	return s.DB.Delete(project).Error
}

func (s *ProjectService) ListUserProjects(userID uint) ([]models.Project, error) {
	// personal projects
	var projects []models.Project
	err := s.DB.Where("user_id = ?", userID).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *ProjectService) ListTeamProjects(teamID uint) ([]models.Project, error) {
	var projects []models.Project
	err := s.DB.Where("team_id = ?", teamID).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
