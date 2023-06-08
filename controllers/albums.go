package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
)

func PostAlbum(context *gin.Context) {
	var albumInput models.CreateAlbumInput
	if err := context.ShouldBindJSON(&albumInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	album := models.Album{
		Title: albumInput.Title,
		Artist: albumInput.Artist,
		Price: albumInput.Price,
	}

	savedAlbum, err := album.Save()
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{ "data": savedAlbum })
}

func GetAlbums(context *gin.Context) {
	var album models.Album
	result := album.FindAll()
	context.IndentedJSON(http.StatusOK, gin.H{ "data": result })
}

func GetAlbumByID(context *gin.Context) {
	var album models.Album
	result, err := album.FindAlbumByID(context.Param("albumID"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found" })
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{ "data": result })
}

func UpdateAlbum(context *gin.Context) {
	var updateAlbumInput models.UpdateAlbumInput
	if err := context.ShouldBindJSON(&updateAlbumInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	var album models.Album
	updatedAlbum, err := album.Update(updateAlbumInput, context.Param("albumID"))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{ "data": updatedAlbum })
}

func DeleteAlbum(context *gin.Context) {
  // Get model if exist
  var album models.Album
	result, err := album.Delete(context.Param("questID"))
  if err != nil {
    context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
    return
  }
  context.IndentedJSON(http.StatusOK, gin.H{ "data": result })
}
