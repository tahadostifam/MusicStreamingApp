package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/dto"
	"github.com/tahadostifam/MusicStreamingApp/api/presenters"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"github.com/tahadostifam/MusicStreamingApp/pkg/jwt_manager"
)

type AuthController struct {
	jwtManager  *jwt_manager.JwtManager
	authService *services.AuthService
}

func NewAuthController(jwtManager *jwt_manager.JwtManager, authService *services.AuthService) *AuthController {
	return &AuthController{jwtManager, authService}
}

func (c *AuthController) HandleSignin(ctx *gin.Context) {
	body := dto.Validate[dto.UserSigninDto](ctx)
	email := body.Email
	password := body.Password

	user, err := c.authService.GetUser(email, password)
	if err != nil {
		presenters.IncorrectPassword(ctx)
	}

	// everything is ok from client side
	c.jwtManager.
		presenters.SendUser(ctx, user)
}

func (c *AuthController) HandleLogout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *AuthController) HandleAuthentication(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
