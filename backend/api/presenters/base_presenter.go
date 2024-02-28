package presenters

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest(ctx *gin.Context) error {
	code := http.StatusBadRequest

	ctx.JSON(code, JsonMessage{
		Code:    code,
		Message: "Bad Request",
	})

	return nil
}
