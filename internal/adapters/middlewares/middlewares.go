package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/internal/adapters/helpers"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := helpers.ExtractToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		if _, err := helpers.ValidateJWToken(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
