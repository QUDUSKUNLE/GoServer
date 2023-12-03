package handlers

import (
	"net/http"
	"os"
	"server/internal/core/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

func (service *HTTPHandler) SaveProfile(ctx *gin.Context) {
	profileDto := domain.ProfileDTO{}
	if err := ctx.ShouldBindJSON(&profileDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false })
		return
	}
	UserID, fal := ctx.Get("UserID")
	if !fal {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "status": false })
		return
	}
	
	user, er := service.ServicesAdapter.ReadUser(UserID.(string))
	if er != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "status": false })
		return
	}

	if err := service.ServicesAdapter.SaveProfile(
		domain.Profile{
			FirstName: strings.TrimSpace(profileDto.FirstName),
			LastName:  strings.TrimSpace(profileDto.LastName),
			UserID: user.ID,
		}); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "Profile created successfully.", "status": true})
}

func (service *HTTPHandler) ReadProfile(ctx *gin.Context) {	
	profile, err := service.ServicesAdapter.ReadProfile(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": profile, "status": true })
}

func (service *HTTPHandler) ReadProfiles(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" || role != os.Getenv("ROLE") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "status": false})
		return
	}
	profiles, err := service.ServicesAdapter.ReadProfiles()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": profiles, "status": true })
}

func (service *HTTPHandler) PatchProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Patch a Profile", "status": true})
}
