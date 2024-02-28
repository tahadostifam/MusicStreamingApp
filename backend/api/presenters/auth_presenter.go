package presenters

import (
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
	"net/http"
)

type UserDto struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	User    *models.User `json:"user"`
}

func IncorrectPassword(ctx *gin.Context) {
	code := http.StatusBadRequest

	ctx.JSON(code, JsonMessage{
		Code:    401,
		Message: services.ErrIncorrectPassword.Error(),
	})
}

func SendUser(ctx *gin.Context, user *models.User) {
	ctx.JSON(http.StatusOK, UserDto{
		Message: "Success",
		Code:    200,
		User:    user,
	})
}
