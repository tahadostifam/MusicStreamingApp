package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tahadostifam/MusicStreamingApp/api/controllers"
)

func InitApi(port int) error {
	app := fiber.New(fiber.Config{})

	authController := controllers.NewAuthController()

	app.Get("/", authController.HandleSignin)

	err := app.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	return nil
}
