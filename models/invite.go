package models

import "time"

type TeamInvite struct {
	ID        uint   `gorm:"primaryKey"`
	TeamID    uint   `gorm:"not null"`
	InviterID uint   `gorm:"not null"`
	InviteeID uint   `gorm:"not null"`
	Status    string // "pending", "accepted", "declined"
	CreatedAt time.Time
	UpdatedAt time.Time
}
