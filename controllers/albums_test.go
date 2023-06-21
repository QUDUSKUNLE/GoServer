package controllers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"
)

func NewUser() *httptest.ResponseRecorder {
	user := models.UserInput{
		Email:    "quduskunle@gmail.co",
		Password: "test",
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

	writer := makeRequest("POST", "/v1/albums", album, true, "quduskunle@gmail.co", "test")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	data, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, data.Title, album.Title)
	assert.Equal(t, data.Artist, album.Artist)
	assert.Equal(t, data.Price, album.Price)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAlbums(t *testing.T) {
	writer := makeRequest("GET", "/v1/albums", nil, true, "quduskunle@gmail.co", "test")
	var response map[string][]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetAlbum(t *testing.T) {
	writer := makeRequest("GET", "/v1/albums/1", nil, true, "quduskunle@gmail.co", "test")
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
	writer := makeRequest("PATCH", "/v1/albums/1", quest, true, "quduskunle@gmail.co", "test")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestDeleteAlbum(t *testing.T) {
	writer := makeRequest("DELETE", "/v1/albums/1", nil, true, "quduskunle@gmail.co", "test")
	var response map[string]models.Album
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}
