package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"server/models"
	"server/controllers"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", controllers.Home)

	// Quests Endpoints
	router.GET("/quests", controllers.FindQuests)
	router.POST("/quests", controllers.CreateQuest)
	router.GET("/quests/:id", controllers.FindQuest)
	router.PATCH("/quests/:id", controllers.UpdateQuest)
	router.DELETE("/quests/:id", controllers.DeleteQuest)

	// Album Endpoints
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbum)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	models.ConnectDatabase()
	if err := router.Run("localhost:"+port); err != nil {
		return
	}
}
