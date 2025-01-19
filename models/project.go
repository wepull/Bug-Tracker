package models

import "time"

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	UserID      *uint // personal project
	TeamID      *uint // team project
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
