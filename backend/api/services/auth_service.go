package services

import (
	"errors"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/auth"
	"github.com/tahadostifam/MusicStreamingApp/pkg/encrypt_password"
)

var ErrIncorrectPassword = errors.New("incorrect password")

type AuthService struct {
	authRepo auth.Repository
}

func NewAuthService(authRepo auth.Repository) *AuthService {
	return &AuthService{authRepo}
}

func (c AuthService) GetUser(email, password string) (*models.User, error) {
	user, err := c.authRepo.FindBy(email)
	if err != nil && user != nil {
		return nil, err
	}

	// validating entered password
	isValid := encrypt_password.CheckPassword(password, user.Password)
	if isValid {
		return user, nil
	}

	return nil, ErrIncorrectPassword
}

func (c AuthService) CreateUser(name, email, password string) (*models.User, error) {
	user, err := c.authRepo.Create(name, email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
