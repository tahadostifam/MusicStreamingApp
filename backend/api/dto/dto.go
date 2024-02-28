package dto

import (
	"github.com/tahadostifam/MusicStreamingApp/api/presenters"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type validationErrorResponse struct {
	Success bool            `json:"success"`
	Errors  []errorResponse `json:"errors"`
}

type errorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

func Validate[Dto interface{}](ctx *gin.Context) *Dto {
	var validationErrors []errorResponse

	body := new(Dto)

	bindErr := ctx.Bind(&body)

	if bindErr != nil {
		presenters.BadRequest(ctx)
		return nil
	}

	errs := validate.Struct(body)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) { //nolint
			var elem errorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, validationErrorResponse{
			Success: false,
			Errors:  validationErrors,
		})

		return nil
	}

	return body
}
