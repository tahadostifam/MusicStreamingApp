package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model `gorm:"primary_key"`
	MusicID    uuid.UUID `gorm:"not null; unique"`
	ArtistID   uuid.UUID `gorm:"not null;"`
	Title      string    `gorm:"not null"`
	Genre      string
	Duration   time.Duration `gorm:"not null"`
	Plays      int           `gorm:"default:0"`

	Comments []Comment `gorm:"many2many:music_comments;joinForeignKey:MusicID"`
}

func (m *Music) BeforeCreate(tx *gorm.DB) (err error) {
	m.MusicID = uuid.New()

	return nil
}
