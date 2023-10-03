package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "server/internal/adapters/handlers"
	repository "server/internal/adapters/repository"
	services "server/internal/core/services"
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
		store := repository.PostgresDatabaseAdapter(
			os.Getenv("HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
		svc = services.ServicesAdapter(store)
	}
	InitializeRoutes()
} 

func InitializeRoutes() {
	port := os.Getenv("PORT")
	router := gin.Default()

	httpHandler := handlers.HTTPAdapter(*svc)
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"mesage": "Welcome to e-Commerce HalalMeat"})
	})
	
	publicRoutes := router.Group("/v1")
	publicRoutes.POST("/users/register", httpHandler.SaveUser)
	publicRoutes.POST("/users/login", httpHandler.Login)

	// ProtectedRoutes Endpoints
	protectedRoutes := router.Group("/v1")
	protectedRoutes.Use(httpHandler.JWTAuthentication())

	// Address Endpoints
	protectedRoutes.POST("/addresses", httpHandler.SaveAddress)
	protectedRoutes.GET("/addresses", httpHandler.ReadAddresses)
	protectedRoutes.GET("/addresses/:id", httpHandler.UUIDMiddleware(), httpHandler.ReadAddress)
	protectedRoutes.PATCH("/addresses/:id", httpHandler.UUIDMiddleware(), httpHandler.PatchAddress)
	protectedRoutes.DELETE("/addresses/:id", httpHandler.UUIDMiddleware(), httpHandler.DeleteAddress)

	// Profile Endpoints
	protectedRoutes.POST("/profiles", httpHandler.SaveProfile)
	protectedRoutes.GET("/profiles", httpHandler.ReadProfiles)
	protectedRoutes.GET("/profiles/:id", httpHandler.UUIDMiddleware(), httpHandler.ReadProfile)
	protectedRoutes.PATCH("/profiles/:id", httpHandler.UUIDMiddleware(), httpHandler.PatchProfile)

	if err := router.Run("localhost:" + port); err != nil {
		return
	}
}
