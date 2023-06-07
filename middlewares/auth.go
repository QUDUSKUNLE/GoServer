package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/helpers"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
			if err := helpers.ValidateJWT(context); err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
				context.Abort()
				return
			}
			context.Next()
	}
}
