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
	var profileInput models.ProfileInput
	var profile models.Profile

	if err := context.ShouldBindJSON(&profileInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}
	profile = models.Profile{
		FirstName: strings.TrimSpace(profileInput.FirstName),
		LastName:  strings.TrimSpace(profileInput.LastName),
		UserID:    user.ID,
	}

	_, err = profile.Save()
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": "Profile submitted successfuly"})
}

func GetProfile(context *gin.Context) {
	var profile models.Profile
	result, err := profile.FindProfile(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetProfiles(context *gin.Context) {
	var role string = context.Query("role")
	if role == "" || role != os.Getenv("ROLE") {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var profile models.Profile
	result := profile.FindProfiles()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func PatchProfile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Patch a Profile"})
}
