package controllers

import (
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"net/http"
	"server/models"
	"testing"
)

func TestAddQuest(t *testing.T) {
	quest := models.Quest{
		Title:       "Arrow of God",
		Description: "Okay",
		Reward:      12,
	}

	writer := makeRequest("POST", "/v1/quests", quest, true, "qudus@gmail.co", "test12345")
	var response map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetQuests(t *testing.T) {
	writer := makeRequest("GET", "/v1/quests", nil, true, "qudus@gmail.co", "test12345")
	var response map[string][]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetQuest(t *testing.T) {
	writer := makeRequest("GET", "/v1/quests/1", nil, true, "qudus@gmail.co", "test12345")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestPatchQuest(t *testing.T) {
	quest := models.Quest{
		Title:       "Arrow of Name",
		Description: "Okay",
		Reward:      12,
	}
	writer := makeRequest("PATCH", "/v1/quests/1", quest, true, "qudus@gmail.co", "test12345")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, false, exists)
	assert.Equal(t, http.StatusNoContent, writer.Code)
}

func TestDeleteQuest(t *testing.T) {
	writer := makeRequest("DELETE", "/v1/quests/1", nil, true, "qudus@gmail.co", "test12345")
	var response map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &response)
	assert.Equal(t, http.StatusNoContent, writer.Code)
}
