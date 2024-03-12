package presenters

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/api/services"
)

type UserDto struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	User    *models.User `json:"user"`
}

func Unauthorized(ctx *gin.Context) {
	code := http.StatusUnauthorized

	ctx.JSON(code, JsonMessage{
		Code:    401,
		Message: "Unauthorized",
	})
}

func IncorrectPassword(ctx *gin.Context) {
	code := http.StatusUnauthorized

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: services.ErrIncorrectPassword.Error(),
	})
}

func SendUser(ctx *gin.Context, user *models.User) {
	code := http.StatusOK
	ctx.JSON(code, UserDto{
		Message: "Success",
		Code:    code,
		User:    user,
	})
}

func EmailAlreadyExist(ctx *gin.Context) {
	code := http.StatusConflict
	ctx.JSON(code, JsonMessage{
		Message: "Another account already exists with the same email address",
		Code:    code,
	})
}
