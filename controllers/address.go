package controllers

import (
	"net/http"
	"server/models"
	"server/middlewares"
	"server/helpers"
	"github.com/gin-gonic/gin"
)

func AddAddress(context *gin.Context) {
	var addressInput models.AddressInput
	var profile models.Profile
	if err := context.ShouldBindJSON(&addressInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": middlewares.CompileErrors(err) })
		return
	}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": middlewares.CompileErrors(err) })
		return
	}
	findProfile, err := profile.FindProfileByUserID((user.ID).String())
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "error": "Profile record not found" })
		return
	}
	address := models.Address{
		StreetNo: addressInput.StreetNo,
		StreetName: addressInput.StreetName,
		Province: addressInput.Province,
		State: addressInput.State,
		ProfileID: findProfile.ID,
	}
	_, err = address.Save()
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{ "error": err.Error() })
		return
	}
	context.JSON(http.StatusCreated, gin.H{ "data": "Address submitted successfully."})
}

func GetAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Get an Address" })
}

func GetAddresses(context *gin.Context) {
	var address models.Address
	result := address.FindAddresses()
	context.JSON(http.StatusOK, gin.H{ "data": result })
}

func PatchAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Patch an Address" })
}

func DeleteAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Delete an Address" })
}
