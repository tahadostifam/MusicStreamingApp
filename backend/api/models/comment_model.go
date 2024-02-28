package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model `gorm:"primary_key"`
	CommentID  string        `gorm:"not null;unique"`
	Content    string        `gorm:"not null"`
	Date       time.Duration `gorm:"not null"`
	Likes      []User        // References to UserID
	UserID     User          `gorm:"not null"` // Foreign key for User
	MusicID    Music         `gorm:"not null"` // Foreign key for Music
}
