package handlers

import (
	"net/http"
	"os"
	"strings"
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
	ctx.JSON(http.StatusCreated, gin.H{"data": "Profile created successfully.", "status": true})
}

func (service *HTTPHandler) ReadProfile(ctx *gin.Context) {	
	profile, err := service.ExternalServicesAdapter.ReadProfile(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": profile })
}

func (service *HTTPHandler) ReadProfiles(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" || role != os.Getenv("ROLE") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	profiles, err := service.ExternalServicesAdapter.ReadProfiles()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": profiles })
}

func (service *HTTPHandler) PatchProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Patch a Profile"})
}
