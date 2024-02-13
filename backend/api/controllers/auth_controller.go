package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c AuthController) HandleSignin(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c AuthController) HandleLogout(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c AuthController) HandleAuthentication(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
