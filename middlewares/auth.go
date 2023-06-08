package middlewares

import (
	"net/http"
	"server/helpers"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := helpers.ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication is required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
