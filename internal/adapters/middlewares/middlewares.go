package middlewares

import (
	"net/http"
	"server/internal/adapters/helpers"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := helpers.ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		// _, err := helpers.CurrentUser(context)
		// if err != nil {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request."})
		// 	context.Abort()
		// 	return
		// }
		context.Next()
	}
}
