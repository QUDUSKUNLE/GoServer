package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func PostAlbums(cnx *gin.Context) {
	var newAlbum models.Album
	if err := cnx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	cnx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbums(cnx *gin.Context) {
	cnx.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(cnx *gin.Context) {
	id := cnx.Param("id")

	for _, album := range albums {
		if album.ID == id {
			cnx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	cnx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func Home(cnx *gin.Context) {
	cnx.IndentedJSON(http.StatusOK, gin.H{"mesage": "Welcome to Go API"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
