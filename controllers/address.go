package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/helpers"
	"server/middlewares"
	"server/models"
)

func AddAddress(context *gin.Context) {
	addressInput := models.AddressModel{}
	
	if err := context.ShouldBindJSON(&addressInput); err != nil {
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
		StreetNo:   addressInput.StreetNo,
		StreetName: addressInput.StreetName,
		Province:   addressInput.Province,
		State:      addressInput.State,
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
	addressModel := models.Address{}
	if _, err := addressModel.Delete(context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "Deleted an address successfully"})
}
