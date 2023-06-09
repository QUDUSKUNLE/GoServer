package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /quests
// GET all quests
func GetQuests(context *gin.Context) {
	var quests []models.Quest
	models.DB.Find(&quests)
	context.JSON(http.StatusOK, gin.H{"data": quests})
}

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

func GetQuest(context *gin.Context) {
	var quest models.Quest

	if err := models.DB.Where("id = ?", context.Param("questID")).First(&quest).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": quest})
}

func UpdateQuest(context *gin.Context) {
	var quest models.Quest
	if err := models.DB.Where("id = ?", context.Param("questID")).First(&quest).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var updateQuestInput models.UpdateQuestInput
	if err := context.ShouldBindJSON(&updateQuestInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&quest).Updates(updateQuestInput)
	context.JSON(http.StatusOK, gin.H{"data": quest})
}

func DeleteQuest(context *gin.Context) {
	// Get model if exist
	var quest models.Quest
	if err := models.DB.Where("id = ?", context.Param("questID")).First(&quest).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&quest)
	context.JSON(http.StatusOK, gin.H{"data": "Deleted successfully"})
}
