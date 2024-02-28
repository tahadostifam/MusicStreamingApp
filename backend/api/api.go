package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/controllers"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/auth"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"github.com/tahadostifam/MusicStreamingApp/config"
	"github.com/tahadostifam/MusicStreamingApp/pkg/jwt_manager"
	"gorm.io/gorm"
)

const JwtDefaultDuration = "120h" // 5 days

func InitApi(host string, port int, db *gorm.DB, config *config.Config) error {
	// initialize repository, services and third-party libraries.
	jwtManager := jwt_manager.NewJwtManager(config.App.Auth.SecretKey, JwtDefaultDuration)
	authRepo := auth.NewRepository(db)
	authService := services.NewAuthService(*authRepo)

	app := gin.Default()

	// initialize controllers
	authController := controllers.NewAuthController(jwtManager, authService)

	app.GET("/", authController.HandleSignin)

	err := app.Run(fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	return nil
}
