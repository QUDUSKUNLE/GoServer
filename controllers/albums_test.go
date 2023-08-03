package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewUser() *httptest.ResponseRecorder {
	user := models.UserInputDto{
		Email:    "qudus@gmail.co",
		Password: "test12345",
	}
	writer := makeRequest("POST", "/v1/users/register", user, false, "", "")
	return writer
}

func TestAddAlbum(t *testing.T) {
	album := models.CreateAlbumInput{
		Title:  "Arrow of God",
		Artist: "Okay",
		Price:  12.45,
	}
	NewUser()

	writer := makeRequest("POST", "/v1/albums", album, true, "qudus@gmail.co", "test12345")
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAlbums(t *testing.T) {
	writer := makeRequest("GET", "/v1/albums", nil, true, "qudus@gmail.co", "test12345")
	var response map[string][]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetAlbum(t *testing.T) {
	writer := makeRequest("GET", "/v1/albums/1", nil, true, "qudus@gmail.co", "test12345")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestPatchAlbum(t *testing.T) {
	quest := models.Album{
		Title:  "Arrow of Name",
		Artist: "Okay",
		Price:  12,
	}
	writer := makeRequest("PATCH", "/v1/albums/1", quest, true, "qudus@gmail.co", "test12345")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestDeleteAlbum(t *testing.T) {
	writer := makeRequest("DELETE", "/v1/albums/1", nil, true, "qudus@gmail.co", "test12345")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}
