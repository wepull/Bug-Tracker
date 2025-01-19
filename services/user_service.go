package services

import (
	"errors"

	"github.com/wepull/Bug-Tracker/dto/requests"
	"github.com/wepull/Bug-Tracker/models"
	"github.com/wepull/Bug-Tracker/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// List all users
func (s *UserService) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Get user by username
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Get user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// Update user details
func (s *UserService) UpdateUser(userID uint, req requests.UpdateUserRequest) (*models.User, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// If updating username, check uniqueness
	if req.Username != nil {
		var existing models.User
		if err := s.DB.Where("username = ?", *req.Username).
			Where("id <> ?", userID).
			First(&existing).Error; err == nil {
			return nil, errors.New("username already taken")
		}
		user.Username = *req.Username
	}

	// If updating email, check uniqueness
	if req.Email != nil {
		var existing models.User
		if err := s.DB.Where("email = ?", *req.Email).
			Where("id <> ?", userID).
			First(&existing).Error; err == nil {
			return nil, errors.New("email already taken")
		}
		user.Email = *req.Email
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	// If updating password, check complexity
	if req.Password != nil {
		if !utils.ValidatePassword(*req.Password) {
			return nil, errors.New("password does not meet complexity requirements")
		}
		hashed, err := utils.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hashed
	}

	if err := s.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete user
func (s *UserService) DeleteUser(userID uint) error {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}
	return s.DB.Delete(user).Error
}
