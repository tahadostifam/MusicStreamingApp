package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model `gorm:"primary_key"`
	Name       string `gorm:"not null"`

	Musics []Music // One-to-Many relationship
}
