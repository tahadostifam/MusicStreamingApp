package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/dto"
	"github.com/tahadostifam/MusicStreamingApp/api/presenters"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/auth"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"github.com/tahadostifam/MusicStreamingApp/pkg/jwt_manager"
)

// AccessToken is the token type of jwt_manager lib
var AccessToken = "token"

type AuthController struct {
	jwtManager  *jwt_manager.JwtManager
	authService *services.AuthService
}

func NewAuthController(jwtManager *jwt_manager.JwtManager, authService *services.AuthService) *AuthController {
	return &AuthController{jwtManager, authService}
}

func (c *AuthController) HandleSignin(ctx *gin.Context) {
	body := dto.Validate[dto.UserSigninDto](ctx)

	if body != nil {
		email := body.Email
		password := body.Password

		user, err := c.authService.GetUser(email, password)
		if err != nil {
			presenters.IncorrectPassword(ctx)
			return
		}

		//everything is ok from client side
		token, tokenErr := c.jwtManager.Generate(AccessToken, user.UserID.String())
		if tokenErr != nil {
			presenters.ServerError(ctx)
			return
		}

		ctx.Header("authentication", token)

		presenters.SendUser(ctx, user)
	}
}

func (c *AuthController) HandleSignup(ctx *gin.Context) {
	body := dto.Validate[dto.UserSignupDto](ctx)

	if body != nil {
		name := body.Name
		email := body.Email
		password := body.Password

		user, err := c.authService.CreateUser(name, email, password)
		if err != nil {
			if errors.Is(err, auth.ErrEmailAlreadyExist) {
				presenters.EmailAlreadyExist(ctx)
			} else {
				presenters.ServerError(ctx)
			}

			return
		}

		// ready to generate access token and deliver it to client
		token, tokenErr := c.jwtManager.Generate(AccessToken, user.UserID.String())
		if tokenErr != nil {
			presenters.ServerError(ctx)
			return
		}

		ctx.Header("authentication", token)

		presenters.SendUser(ctx, user)
	}
}

func (c *AuthController) HandleLogout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *AuthController) HandleAuthentication(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
