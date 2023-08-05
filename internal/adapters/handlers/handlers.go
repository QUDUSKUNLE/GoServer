package handlers

import (
	services "server/internal/core/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"server/internal/adapters/helpers"
)

type HTTPHandler struct {
	ExternalServicesAdapter  services.ServicesHandler
	InternalServicesAdapter  services.ServicesHandler
}

func HTTPAdapter(services services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ExternalServicesAdapter: services,
		InternalServicesAdapter: services,
	}
}

func (service *HTTPHandler) JWTAuthMiddleware() gin.HandlerFunc {
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

func (service *HTTPHandler) UUidMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !helpers.ValidateUUID(context.Param("id")) {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id credential"})
			return
		}
		context.Next()
	}
}
