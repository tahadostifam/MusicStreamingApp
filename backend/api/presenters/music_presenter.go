package presenters

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
)

type MusicDto struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Music   *models.Music `json:"music"`
}

func MusicFileRequired(ctx *gin.Context) {
	code := http.StatusBadRequest

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "MusicFile required",
	})
}

func MaxFileUploadSize(ctx *gin.Context, max string) {
	code := http.StatusBadRequest

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "Maxium file upload size is " + max,
	})
}

func MusicAdded(ctx *gin.Context, music *models.Music) {
	code := http.StatusOK

	ctx.JSON(code, MusicDto{
		Code:    code,
		Message: "Music added",
		Music:   music,
	})
}

func MusicDeleted(ctx *gin.Context) {
	code := http.StatusOK

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "Music deleted",
	})
}

func ErrOnlyUserCan(ctx *gin.Context) {
	code := http.StatusForbidden

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: services.ErrOnlyUserCan.Error(),
	})
}

func MusicNotFound(ctx *gin.Context) {
	code := http.StatusNotFound

	ctx.String(code, "404 - Music file not found")
}
