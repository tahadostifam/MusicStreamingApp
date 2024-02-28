package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Music struct {
	gorm.Model `gorm:"primary_key"`
	MusicID    uuid.UUID `gorm:"not null; unique"`
	Title      string    `gorm:"not null"`
	Genre      string
	Duration   time.Duration `gorm:"not null"`
	Plays      int           `gorm:"default:0"`
	Likes      []User        // References to UserID
	UserID     User          `gorm:"not null"` // Foreign key for User
	ArtistID   User          // Foreign key for Artist (Many-to-One)
	Comments   []Comment     // One-to-Many relationship
}
