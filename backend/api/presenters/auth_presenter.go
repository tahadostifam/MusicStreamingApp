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

func EmailAlreadyExist(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, JsonMessage{
		Message: "Another account already exists with the same email address",
		Code:    409, // conflict
	})
}
