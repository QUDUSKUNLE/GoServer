package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddProfile(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{ "data": "Profile submitted successfully."})
}

func GetProfile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Get a Profile" })
}

func GetProfiles(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Get Profiles" })
}

func PatchProfile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Patch a Profile" })
}

func DeleteProfile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Delete a Profile" })
}
