package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model `gorm:"primary_key"`
	CommentID  string        `gorm:"not null;unique"`
	Content    string        `gorm:"not null"`
	Date       time.Duration `gorm:"not null"`
	UserID     string        `gorm:"not null"` // Foreign key for User
	MusicID    string        `gorm:"not null"` // Foreign key for Music
	// Likes      []User        // References to UserID
}
