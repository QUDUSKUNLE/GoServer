package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddQuest(t *testing.T) {
	quest := models.Quest{
		Title: "Arrow of God",
		Description: "Okay",
		Reward: 12,
	}

	writer := makeRequest("POST", "/api/quests", quest, true, "quduskunle", "test")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	data, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, data.Title, quest.Title)
	assert.Equal(t, data.Description, quest.Description)
	assert.Equal(t, data.Reward, quest.Reward)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetQuests(t *testing.T) {
	writer := makeRequest("GET", "/api/quests", nil, true, "quduskunle", "test")
	var response map[string][]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetQuest(t *testing.T) {
	writer := makeRequest("GET", "/api/quests/1", nil, true, "quduskunle", "test")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestPatchQuest(t *testing.T) {
	quest := models.Quest{
		Title: "Arrow of Name",
		Description: "Okay",
		Reward: 12,
	}
	writer := makeRequest("PATCH", "/api/quests/1", quest, true, "quduskunle", "test")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestDeleteQuest(t *testing.T) {
	writer := makeRequest("DELETE", "/api/quests/1", nil, true, "quduskunle", "test")
	var response map[string]models.Quest
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
}
