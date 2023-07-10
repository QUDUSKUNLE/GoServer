package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddAlbum(context *gin.Context) {
	albumInput := models.CreateAlbumInput{}
	if err := context.ShouldBindJSON(&albumInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album := models.Album{
		Title:  albumInput.Title,
		Artist: albumInput.Artist,
		Price:  albumInput.Price,
	}

	savedAlbum, err := album.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedAlbum})
}

func GetAlbums(context *gin.Context) {
	album := models.Album{}
	result := album.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetAlbum(context *gin.Context) {
	album := models.Album{}
	result, err := album.FindAlbumByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateAlbum(context *gin.Context) {
	updateAlbumInput := models.UpdateAlbumInput{}
	if err := context.ShouldBindJSON(&updateAlbumInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{}
	updatedAlbum, err := album.Update(updateAlbumInput, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedAlbum})
}

func DeleteAlbum(context *gin.Context) {
	// Get model if exist
	album := models.Album{}
	result, err := album.Delete(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}
