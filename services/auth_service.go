package services

import (
	"errors"
	"fmt"

	"github.com/wepull/Bug-Tracker/config"
	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/models"
	"github.com/wepull/Bug-Tracker/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB     *gorm.DB
	Config *config.Config
}

func (s *AuthService) Register(req requests.RegisterRequest) (*models.User, error) {
	// 1. Validate password complexity
	if !utils.ValidatePassword(req.Password) {
		return nil, errors.New("password must be >=8 chars, with upper, lower, digit, symbol")
	}

	// 2. Check uniqueness
	var existing models.User
	if err := s.DB.Where("username = ? OR email = ?", req.Username, req.Email).
		First(&existing).Error; err == nil {
		return nil, errors.New("username or email already in use")
	}

	// 3. Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 4. Create User
	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashed,
	}
	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}

	// 5. Create default team
	defaultTeam := models.Team{
		Name:      fmt.Sprintf("%s_default_team", user.Username),
		CreatedBy: user.ID,
	}
	if err := s.DB.Create(&defaultTeam).Error; err != nil {
		return nil, err
	}

	// Add membership
	teamMember := models.TeamMember{
		UserID: user.ID,
		TeamID: defaultTeam.ID,
	}
	if err := s.DB.Create(&teamMember).Error; err != nil {
		return nil, err
	}

	// 6. Create default project (associated with that team or personal)
	defaultProject := models.Project{
		Name:   "Default Project",
		TeamID: &defaultTeam.ID,
	}
	if err := s.DB.Create(&defaultProject).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(usernameOrEmail, password string) (*models.User, error) {
	var user models.User
	// Attempt to find by username first
	err := s.DB.Where("username = ?", usernameOrEmail).First(&user).Error
	if err != nil {
		// If not found, try email
		errEmail := s.DB.Where("email = ?", usernameOrEmail).First(&user).Error
		if errEmail != nil {
			return nil, errors.New("invalid username/email or password")
		}
	}

	// Check password
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid username/email or password")
	}

	return &user, nil
}
