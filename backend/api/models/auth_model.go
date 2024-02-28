package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `gorm:"primary_key"`
	UserID     uuid.UUID `gorm:"not null;unique"`
	Name       string
	Email      string `gorm:"unique"`
	Password   string `gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.New()

	return nil
}
