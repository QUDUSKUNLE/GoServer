package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddAddress(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{ "data": "Address submitted successfully."})
}

func GetAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Get an Address" })
}

func GetAddresses(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Get Addresses" })
}

func PatchAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Patch an Address" })
}

func DeleteAddress(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "Delete an Address" })
}
