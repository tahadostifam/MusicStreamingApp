package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model `gorm:"primary_key"`
	MusicID    string        `gorm:"not null; unique"`
	ArtistID   string        `gorm:"not null;"`
	Title      string        `gorm:"not null"`
	Genre      string        `gorm:"not null"`
	Duration   time.Duration `gorm:"not null"`
	FileName   string        `gorm:"not null"`
	Plays      int           `gorm:"default:0"`

	Comments []Comment `gorm:"many2many:music_comments;joinForeignKey:MusicID"`
}

func (m *Music) BeforeCreate(tx *gorm.DB) (err error) {
	m.MusicID = uuid.NewString()

	return nil
}
