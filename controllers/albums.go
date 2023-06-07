package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

func PostAlbum(cnx *gin.Context) {
	var albumInput models.CreateAlbumInput
	if err := cnx.ShouldBindJSON(&albumInput); err != nil {
		cnx.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	album := models.Album{ Title: albumInput.Title, Artist: albumInput.Artist, Price: albumInput.Price }
	models.DB.Create(&album)

	cnx.IndentedJSON(http.StatusCreated, gin.H{ "data": album })
}

func GetAlbums(cnx *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	cnx.IndentedJSON(http.StatusOK, gin.H{ "data": albums })
}

func GetAlbumByID(cnx *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ? ", cnx.Param("id")).First(&album).Error; err != nil {
		cnx.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found" })
		return
	}
	cnx.IndentedJSON(http.StatusOK, gin.H{ "data": album })
}

func UpdateAlbum(cnx *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", cnx.Param("id")).First(&album).Error; err != nil {
		cnx.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	var updateAlbumInput models.UpdateAlbumInput
	if err := cnx.ShouldBindJSON(&updateAlbumInput); err != nil {
		cnx.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	models.DB.Model(&album).Updates(updateAlbumInput)
	cnx.IndentedJSON(http.StatusOK, gin.H{ "data": album })
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


func Home(cnx *gin.Context) {
	cnx.IndentedJSON(http.StatusOK, gin.H{ "mesage": "Welcome to Go API" })
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
