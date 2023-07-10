package controllers

import (
	"net/http"
	"os"
	"server/helpers"
	"server/middlewares"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddProfile(context *gin.Context) {
	profileInputModel := models.ProfileInput{}
	if err := context.ShouldBindJSON(&profileInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	profile := models.Profile{
		FirstName: strings.TrimSpace(profileInputModel.FirstName),
		LastName:  strings.TrimSpace(profileInputModel.LastName),
		UserID:    user.ID,
	}

	if err := profile.Save(); err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": "Profile submitted successfuly"})
}

func GetProfile(context *gin.Context) {
	profileModel := models.Profile{}
	
	profile, err := profileModel.FindProfile(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": profile})
}

func GetProfiles(context *gin.Context) {
	role := context.Query("role")
	if role == "" || role != os.Getenv("ROLE") {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	profileModel := models.Profile{}
	profile := profileModel.FindProfiles()
	context.JSON(http.StatusOK, gin.H{"data": profile})
}

func PatchProfile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Patch a Profile"})
}
