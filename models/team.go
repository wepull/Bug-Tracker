package models

import "time"

type Team struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string
	CreatedBy   uint // user ID who created this team
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Many-to-many table
type TeamMember struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	TeamID uint `gorm:"not null"`
}
