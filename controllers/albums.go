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
	var albums []models.Album
	models.DB.Find(&albums)
	context.IndentedJSON(http.StatusOK, gin.H{ "data": albums })
}

func GetAlbumByID(context *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ? ", context.Param("id")).First(&album).Error; err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found" })
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{ "data": album })
}

func UpdateAlbum(context *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", context.Param("id")).First(&album).Error; err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	var updateAlbumInput models.UpdateAlbumInput
	if err := context.ShouldBindJSON(&updateAlbumInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	models.DB.Model(&album).Updates(updateAlbumInput)
	context.IndentedJSON(http.StatusOK, gin.H{ "data": album })
}

func DeleteAlbum(con *gin.Context) {
  // Get model if exist
  var album models.Album
  if err := models.DB.Where("id = ?", con.Param("id")).First(&album).Error; err != nil {
    con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
    return
  }

  models.DB.Delete(&album)
  con.IndentedJSON(http.StatusOK, gin.H{ "data": true })
}
