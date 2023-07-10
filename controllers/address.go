package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/helpers"
	"server/middlewares"
	"server/models"
)

func AddAddress(context *gin.Context) {
	addressInputModel := models.AddressInput{}
	
	if err := context.ShouldBindJSON(&addressInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	profileModel := models.Profile{}

	findProfile, err := profileModel.FindProfileByUserID((user.ID).String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Profile record not found"})
		return
	}

	addressModel := models.Address{
		StreetNo:   addressInputModel.StreetNo,
		StreetName: addressInputModel.StreetName,
		Province:   addressInputModel.Province,
		State:      addressInputModel.State,
		ProfileID:  findProfile.ID,
	}

	if err := addressModel.Save(); err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": "Address submitted successfully."})
}

func GetAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Get an Address"})
}

func GetAddresses(context *gin.Context) {
	addressModel := models.Address{}
	addresses := addressModel.FindAddresses()
	context.JSON(http.StatusOK, gin.H{"data": addresses})
}

func PatchAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Patch an Address"})
}

func DeleteAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Delete an Address"})
}
