package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model `gorm:"primary_key"`
	CommentID  string        `gorm:"not null;unique"`
	Content    string        `gorm:"not null"`
	Date       time.Duration `gorm:"not null"`
	UserID     uuid.UUID     `gorm:"not null"` // Foreign key for User
	MusicID    uuid.UUID     `gorm:"not null"` // Foreign key for Music
	// Likes      []User        // References to UserID
}
