package services

import (
	"errors"

	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/models"
	"gorm.io/gorm"
)

type TeamService struct {
	DB *gorm.DB
}

// Create a new team
func (s *TeamService) CreateTeam(req requests.CreateTeamRequest, creatorID uint) (*models.Team, error) {
	team := &models.Team{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   creatorID,
	}
	if err := s.DB.Create(team).Error; err != nil {
		return nil, err
	}

	// Add membership
	teamMember := models.TeamMember{
		UserID: creatorID,
		TeamID: team.ID,
	}
	if err := s.DB.Create(&teamMember).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// Get team by ID
func (s *TeamService) GetTeamByID(teamID uint) (*models.Team, error) {
	var team models.Team
	if err := s.DB.First(&team, teamID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return &team, nil
}

// Update a team
func (s *TeamService) UpdateTeam(teamID uint, req requests.UpdateTeamRequest) (*models.Team, error) {
	team, err := s.GetTeamByID(teamID)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		team.Name = *req.Name
	}
	if req.Description != nil {
		team.Description = *req.Description
	}
	if err := s.DB.Save(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// Delete a team
func (s *TeamService) DeleteTeam(teamID uint) error {
	team, err := s.GetTeamByID(teamID)
	if err != nil {
		return err
	}
	return s.DB.Delete(team).Error
}

// List teams of a user
func (s *TeamService) ListUserTeams(userID uint) ([]models.Team, error) {
	var teams []models.Team
	err := s.DB.Joins("JOIN team_members ON team_members.team_id = teams.id").
		Where("team_members.user_id = ?", userID).
		Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// Leave a team
func (s *TeamService) RemoveUserFromTeam(teamID, userID uint) error {
	return s.DB.Where("team_id = ? AND user_id = ?", teamID, userID).
		Delete(&models.TeamMember{}).Error
}
