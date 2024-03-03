package presenters

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func BadRequest(ctx *gin.Context) error {
	code := http.StatusBadRequest

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "Bad Request",
	})

	return nil
}

func ServerError(ctx *gin.Context) error {
	code := http.StatusInternalServerError

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "Internal Server Error",
	})

	return nil
}

func DetailedServerError(ctx *gin.Context, err error) error {
	code := http.StatusInternalServerError

	ctx.JSON(code, ErrorMessage{
		Code:    code,
		Message: "Internal Server Error",
		Detail:  err.Error(),
	})

	return nil
}
