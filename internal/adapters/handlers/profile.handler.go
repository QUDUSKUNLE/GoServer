package handlers

import (
	"net/http"
	"strings"
	// "server/internal/adapters/helpers"
	"server/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (service *HTTPHandler) SaveProfile(ctx *gin.Context) {
	profileDto := domain.ProfileInputDto{}
	if err := ctx.ShouldBindJSON(&profileDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}
	user, err := service.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}
	if err := service.ExternalServicesAdapter.SaveProfile(
		domain.Profile{
			FirstName: strings.TrimSpace(profileDto.FirstName),
			LastName:  strings.TrimSpace(profileDto.LastName),
			UserID:    user.ID,
		}); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "Profile created successfully"})
}
