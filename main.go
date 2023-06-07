package main

import (
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"server/models"
	"server/controllers"
	"server/middlewares"
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

	router.GET("/", func (context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{ "mesage": "Welcome to Go API" })
	})

	// Users Endpoints

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	// Quests Endpoints
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.GET("/quests", controllers.FindQuests)
	protectedRoutes.POST("/quests", controllers.CreateQuest)
	protectedRoutes.GET("/quests/:id", controllers.FindQuest)
	protectedRoutes.PATCH("/quests/:id", controllers.UpdateQuest)
	protectedRoutes.DELETE("/quests/:id", controllers.DeleteQuest)

	// Album Endpoints
	protectedRoutes.GET("/albums", controllers.GetAlbums)
	protectedRoutes.POST("/albums", controllers.PostAlbum)
	protectedRoutes.GET("/albums/:id", controllers.GetAlbumByID)
	protectedRoutes.PATCH("/albums/:id", controllers.UpdateAlbum)
	protectedRoutes.DELETE("/albums/:id", controllers.DeleteAlbum)

	models.ConnectDatabase()
	if err := router.Run("localhost:"+port); err != nil {
		return
	}
}
