package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/MusicStreamingApp/api/controllers"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/api/presenters"
)

func AuthRequired(authController controllers.AuthController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Here i just call the authentication action on the controller and everything will be handled there, nothing more.
		authController.Authenticate(ctx, func(user *models.User) {
			if user == nil {
				presenters.Unauthorized(ctx)
				return
			}

			ctx.Set("user_id", user.UserID)
		})
	}
}
