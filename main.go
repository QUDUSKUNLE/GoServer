package main

import (
	"log"
	"net/http"
	"os"
	"server/controllers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"mesage": "Welcome to Go API"})
	})

	// PublicRoutes Endpoints
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	// ProtectedRoutes Endpoints
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())

	// Quest Endpoints
	protectedRoutes.GET("/quests", controllers.GetQuests)
	protectedRoutes.POST("/quests", controllers.AddQuest)
	protectedRoutes.GET("/quests/:questID", controllers.GetQuest)
	protectedRoutes.PATCH("/quests/:questID", controllers.UpdateQuest)
	protectedRoutes.DELETE("/quests/:questID", controllers.DeleteQuest)

	// Album Endpoints
	protectedRoutes.GET("/albums", controllers.GetAlbums)
	protectedRoutes.POST("/albums", controllers.AddAlbum)
	protectedRoutes.GET("/albums/:albumID", controllers.GetAlbum)
	protectedRoutes.PATCH("/albums/:albumID", controllers.UpdateAlbum)
	protectedRoutes.DELETE("/albums/:albumID", controllers.DeleteAlbum)

	models.ConnectDatabase()
	if err := router.Run("localhost:" + port); err != nil {
		return
	}
}
