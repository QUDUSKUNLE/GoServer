package controllers

import (
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
	assert.Equal(t, http.StatusCreated, writer.Code)
}
