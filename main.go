package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"server/controllers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"server/internal/adapters/handlers"
	"server/internal/adapters/repository"
	"server/internal/core/services"
)

var (
	repo = flag.String("db", "postgres", "Database for storing messages")
	svc *services.ServicesHandler
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	flag.Parse()

	switch *repo {
	case "redis":
	default:
		store := repository.NewPostgresDatabase(
			os.Getenv("HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
		svc = services.NewServicesHandler(*store)
	}
	InitializeRoutes()
}

func InitializeRoutes() {
	port := os.Getenv("PORT")
	router := gin.Default()

	handler := handlers.NewHTTPHandlers(*svc)
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"mesage": "Welcome to e-Commerce HalalMeat"})
	})
	
	publicRoutes := router.Group("/v1")
	publicRoutes.POST("/users", handler.SaveUser)
	publicRoutes.POST("/users/register", controllers.Register)
	publicRoutes.POST("/users/login", controllers.Login)

	// ProtectedRoutes Endpoints
	protectedRoutes := router.Group("/v1")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())

	// Quest Endpoints
	protectedRoutes.GET("/quests", controllers.GetQuests)
	protectedRoutes.POST("/quests", controllers.AddQuest)
	protectedRoutes.GET("/quests/:id", middlewares.UUidMiddleware(), controllers.GetQuest)
	protectedRoutes.PATCH("/quests/:id", middlewares.UUidMiddleware(), controllers.UpdateQuest)
	protectedRoutes.DELETE("/quests/:id", middlewares.UUidMiddleware(), controllers.DeleteQuest)

	// Album Endpoints
	protectedRoutes.GET("/albums", controllers.GetAlbums)
	protectedRoutes.POST("/albums", controllers.AddAlbum)
	protectedRoutes.GET("/albums/:id", middlewares.UUidMiddleware(), controllers.GetAlbum)
	protectedRoutes.PATCH("/albums/:id", middlewares.UUidMiddleware(), controllers.UpdateAlbum)
	protectedRoutes.DELETE("/albums/:id", middlewares.UUidMiddleware(), controllers.DeleteAlbum)

	// Stock Endpoints
	protectedRoutes.GET("/stocks", controllers.GetStocks)
	protectedRoutes.POST("/stocks", controllers.AddStock)
	protectedRoutes.GET("/stocks/:id", middlewares.UUidMiddleware(), controllers.GetStock)
	protectedRoutes.PATCH("/stocks/:id", middlewares.UUidMiddleware(), controllers.UpdateStock)
	protectedRoutes.DELETE("/stocks/:id", middlewares.UUidMiddleware(), controllers.DeleteStock)

	// Order Endpoints
	protectedRoutes.POST("/orders", controllers.AddOrder)
	protectedRoutes.GET("/orders", controllers.GetOrders)
	protectedRoutes.GET("/orders/:id", middlewares.UUidMiddleware(), controllers.GetOrder)
	protectedRoutes.PATCH("/orders/:id", middlewares.UUidMiddleware(), controllers.PatchOrder)
	protectedRoutes.DELETE("/orders/:id", middlewares.UUidMiddleware(), controllers.DeleteOrder)

	// Address Endpoints
	protectedRoutes.POST("/addresses", controllers.AddAddress)
	protectedRoutes.GET("/addresses", controllers.GetAddresses)
	protectedRoutes.GET("/addresses/:id", middlewares.UUidMiddleware(), controllers.GetAddress)
	protectedRoutes.PATCH("/addresses/:id", middlewares.UUidMiddleware(), controllers.PatchAddress)
	protectedRoutes.DELETE("/addresses/:id", middlewares.UUidMiddleware(), controllers.DeleteAddress)

	// Profile Endpoints
	protectedRoutes.POST("/profiles", controllers.AddProfile)
	protectedRoutes.GET("/profiles", controllers.GetProfiles)
	protectedRoutes.GET("/profiles/:id", middlewares.UUidMiddleware(), controllers.GetProfile)
	protectedRoutes.PATCH("/profiles/:id", middlewares.UUidMiddleware(), controllers.PatchProfile)

	models.ConnectDatabase()
	if err := router.Run("localhost:" + port); err != nil {
		return
	}
}
