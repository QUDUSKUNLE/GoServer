package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}

func postAlbums(cnx *gin.Context) {
	var newAlbum album
	if err := cnx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	cnx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(cnx *gin.Context) {
	cnx.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(cnx *gin.Context) {
	id := cnx.Param("id")

	res, _ := HashPassword(id)
	fmt.Println(res)

	for _, album := range albums {
		if album.ID == id {
			cnx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	cnx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func home(cnx *gin.Context) {
	cnx.IndentedJSON(http.StatusOK, "Home")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
