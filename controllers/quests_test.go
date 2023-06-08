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

	writer := makeRequest("POST", "/api/quests", quest, true)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["data"]
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusCreated, writer.Code)
}
