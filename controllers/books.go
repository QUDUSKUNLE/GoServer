package controllers


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /books
// GET all books
func FindBooks(cg *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	cg.JSON(http.StatusOK, gin.H{"data": books})
}
