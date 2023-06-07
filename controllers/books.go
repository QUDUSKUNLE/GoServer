package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
)


// GET /books
// GET all books
func FindBooks(cg *gin.Context) {
	var books []models.Quest
	models.DB.Find(&books)
	cg.IndentedJSON(http.StatusOK, gin.H{"data": books})
}
