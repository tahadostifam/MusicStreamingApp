package api

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/controllers"
	"github.com/tahadostifam/MusicStreamingApp/api/middlewares"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/auth"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/music"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"github.com/tahadostifam/MusicStreamingApp/config"
	"github.com/tahadostifam/MusicStreamingApp/pkg/jwt_manager"
	"github.com/tahadostifam/MusicStreamingApp/pkg/minio_client"
	"gorm.io/gorm"
)

const JwtDefaultDuration = "120h" // 5 days

func InitApi(host string, port int, db *gorm.DB, config *config.Config) error {
	// initialize minio-client
	minioClient := minio_client.NewMinioClient(config.MinIO.Endpoint, config.MinIO.AccessKey, config.MinIO.SecretKey)

	// initialize repository, services and third-party libraries.
	jwtManager := jwt_manager.NewJwtManager(config.App.Auth.SecretKey, JwtDefaultDuration)
	authRepo := auth.NewRepository(db)
	authService := services.NewAuthService(*authRepo)

	musicRepo := music.NewRepository(db)
	musicService := services.NewMusicService(*musicRepo, *authRepo, *minioClient)

	app := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())

	// Maximum file size for file upload
	app.MaxMultipartMemory = services.MaxFileUploadSize

	// initialize controllers
	authController := controllers.NewAuthController(jwtManager, authService)
	musicController := controllers.NewMusicController(musicService)

	// init auth-required group
	authRequiredGroup := app.Group("/")
	authRequiredGroup.Use(middlewares.AuthRequired(*authController))

	// define routes
	{
		app.POST("/signin", authController.HandleSignin)
		app.POST("/signup", authController.HandleSignup)
		app.POST("/authentication", authController.HandleAuthentication)
	}
	{
		authRequiredGroup.GET("/music/:file_name", musicController.HandleGetMusic)
		authRequiredGroup.POST("/music/new", musicController.HandleNewMusic)
		authRequiredGroup.DELETE("/music/:file_name", musicController.HandleDeleteMusic)
	}

	// Setup logs
	logsFile, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(logsFile)

	err := app.Run(fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	return nil
}
