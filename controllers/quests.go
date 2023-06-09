package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddQuest(context *gin.Context) {
	// Validate quest
	var questInput models.CreateQuestInput
	if err := context.ShouldBindJSON(&questInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create quest
	quest := models.Quest{
		Title:       questInput.Title,
		Description: questInput.Description,
		Reward:      questInput.Reward,
	}
	savedQuest, err := quest.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedQuest})
}

// GET /quests
// GET all quests
func GetQuests(context *gin.Context) {
	var quest models.Quest
	result := quest.GetQuests()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetQuest(context *gin.Context) {
	var quest models.Quest
	result, err := quest.GetQuest(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateQuest(context *gin.Context) {
	var updateQuestInput models.UpdateQuestInput
	if err := context.ShouldBindJSON(&updateQuestInput); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var quest models.Quest
	updatedQuest, err := quest.UpdateQuest(updateQuestInput, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{"data": updatedQuest})
}

func DeleteQuest(context *gin.Context) {
	// Get model if exist
	var quest models.Quest
	_, err := quest.DeleteQuest(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{"data": "Deleted successfully"})
}
