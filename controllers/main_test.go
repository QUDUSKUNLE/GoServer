package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"server/middlewares"
	"server/models"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setupDB()
	exitCode := m.Run()
	tearDownDB()

	os.Exit(exitCode)
}

func setupDB() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}
	models.ConnectDatabase()
}

func tearDownDB() {
	migrator := models.DB.Migrator()
	migrator.DropTable(&models.User{}, &models.Quest{}, &models.Album{})
}

func apiRouter() *gin.Engine {
	router := gin.Default()

	publicRouters := router.Group("/auth")
	publicRouters.POST("/register", Register)
	publicRouters.POST("/login", Login)

	protectedRouters := router.Group("/api")
	protectedRouters.Use(middlewares.JWTAuthMiddleware())
	protectedRouters.POST("/quests", AddQuest)
	protectedRouters.GET("/quests", GetQuests)
	protectedRouters.GET("/quests/:id", GetQuest)
	protectedRouters.PATCH("/quests/:id", UpdateQuest)
	protectedRouters.DELETE("/quests/:id", DeleteQuest)

	protectedRouters.POST("/albums", AddAlbum)
	protectedRouters.GET("/albums", GetAlbums)
	protectedRouters.GET("/albums/:id", GetAlbum)
	protectedRouters.PATCH("/albums/:id", UpdateAlbum)
	protectedRouters.DELETE("/albums/:id", DeleteAlbum)

	return router
}

func makeRequest(method, url string, body interface{}, isAuthenticated bool, username, password string) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticated {
		request.Header.Add("Authorization", "Bearer "+bearerToken(username, password))
	}
	writer := httptest.NewRecorder()
	apiRouter().ServeHTTP(writer, request)
	return writer
}

func bearerToken(username, password string) string {
	user := models.UserInput{
		Email: username,
		Password: password,
	}

	writer := makeRequest("POST", "/auth/login", user, false, "", "")
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["token"]
}
