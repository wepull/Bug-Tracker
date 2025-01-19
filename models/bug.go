package models

import "time"

type Bug struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Severity    string // "critical", "high", "medium", "low"
	Status      string // "open", "in_progress", "resolved", "closed"
	ProjectID   uint   `gorm:"not null"`
	CreatedBy   uint   `gorm:"not null"`
	AssignedTo  uint   // can be 0 if unassigned
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
