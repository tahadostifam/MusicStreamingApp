package controllers

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/dto"
	"github.com/tahadostifam/MusicStreamingApp/api/presenters"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"github.com/tahadostifam/MusicStreamingApp/config"
	"github.com/tahadostifam/MusicStreamingApp/pkg/random"
)

var MusicFilesTmp = config.ProjectRootPath + "tmp/uploads"

const RandomFileNameLength = 25
const MusicFileKey = "music_file"

type MusicController struct {
	musicService *services.MusicService
}

func NewMusicController(musicService *services.MusicService) *MusicController {
	return &MusicController{musicService}
}

func (c *MusicController) HandleNewMusic(ctx *gin.Context) {
	body := dto.Validate[dto.NewMusicDto](ctx)

	if body != nil {
		// Check the MusicFile
		musicFile, formErr := ctx.FormFile("MusicFile")
		if formErr != nil {
			presenters.MusicFileRequired(ctx)
			return
		}

		if musicFile.Size > services.MaxFileUploadSize {
			presenters.MaxFileUploadSize(ctx, fmt.Sprintf("%vMib", services.MaxFileUploadSizeMegabyte))
			return
		}

		randomFileName := random.GenerateRandomFileName(RandomFileNameLength)
		filePath := fmt.Sprintf("%s/%s", MusicFilesTmp, randomFileName)
		saveErr := ctx.SaveUploadedFile(musicFile, filePath)
		if saveErr != nil {
			presenters.ServerError(ctx)
		}

		artistID := ctx.GetString("user_id")
		if len(artistID) == 0 {
			presenters.Unauthorized(ctx)
			return
		}

		music, err := c.musicService.NewMusic(filePath, artistID, body.Title, body.Genre)
		if err != nil {
			presenters.ServerError(ctx)
			return
		}

		presenters.MusicAdded(ctx, music)

		os.Remove(filePath)
	}
}

func (c *MusicController) HandleGetMusic(ctx *gin.Context) {
	fileName := ctx.Param("file_name")
	if len(fileName) == 0 {
		presenters.BadRequest(ctx)
		return
	}

	filePath, getErr := c.musicService.GetMusic(fileName)
	if getErr != nil {
		presenters.MusicNotFound(ctx)
		return
	}

	ctx.File(filePath)

	os.Remove(filePath)
}

func (c *MusicController) HandleDeleteMusic(ctx *gin.Context) {
	body := dto.Validate[dto.DeleteMusicDto](ctx)

	if body != nil {
		artistID := ctx.GetString("user_id")
		err := c.musicService.DeleteMusic(artistID, body.MusicID)
		if errors.Is(err, services.ErrOnlyUserCan) {
			presenters.ErrOnlyUserCan(ctx)
			return
		} else if errors.Is(err, services.ErrMusicFileNotFound) {
			presenters.MusicNotFound(ctx)
			return
		} else if err != nil {
			presenters.ServerError(ctx)
			return
		}

		presenters.MusicDeleted(ctx)
	}
}
