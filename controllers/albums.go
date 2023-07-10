package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddAlbum(context *gin.Context) {
	albumInputModel := models.CreateAlbumInput{}

	if err := context.ShouldBindJSON(&albumInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	albumModel := models.Album{
		Title:  albumInputModel.Title,
		Artist: albumInputModel.Artist,
		Price:  albumInputModel.Price,
	}

	if err := albumModel.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": "Album created successfully"})
}

func GetAlbums(context *gin.Context) {
	album := models.Album{}
	result := album.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetAlbum(context *gin.Context) {
	albumModel := models.Album{}
	result, err := albumModel.FindAlbumByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateAlbum(context *gin.Context) {
	updateAlbumInputModel := models.UpdateAlbumInput{}
	if err := context.ShouldBindJSON(&updateAlbumInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{}
	updatedAlbum, err := album.Update(updateAlbumInputModel, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedAlbum})
}

func DeleteAlbum(context *gin.Context) {
	// Get model if exist
	albumModel := models.Album{}

	result, err := albumModel.Delete(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}
