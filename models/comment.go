package models

import "time"

type Comment struct {
	ID          uint   `gorm:"primaryKey"`
	BugID       uint   `gorm:"not null"`
	AuthorID    uint   `gorm:"not null"`
	CommentText string `gorm:"type:text"`
	CreatedAt   time.Time
}
