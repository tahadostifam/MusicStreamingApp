package services

import "github.com/tahadostifam/MusicStreamingApp/api/models"

type AuthService struct{}

func NewAuthController() *AuthService {
	return &AuthService{}
}

func (c AuthService) HandleSignin(email string) (err error, code int) {
	//TODO implement me
	panic("implement me")
}

func (c AuthService) HandleLogout(token string) error {
	//TODO implement me
	panic("implement me")
}

func (c AuthService) HandleAuthentication(token string) (error, models.UserModel) {
	//TODO implement me
	panic("implement me")
}
