package models

import (
	"github.com/google/uuid"
	"github.com/tahadostifam/MusicStreamingApp/pkg/encrypt_password"
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
	u.Password = encrypt_password.EncryptPassword(u.Password)

	return nil
}
