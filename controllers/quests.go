package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
)

// GET /quests
// GET all quests
func FindQuests(context *gin.Context) {
	var quests []models.Quest
	models.DB.Find(&quests)
	context.IndentedJSON(http.StatusOK, gin.H{ "data": quests })
}

func CreateQuest(context *gin.Context) {
	// Validate quest
	var questInput models.CreateQuestInput
	if err := context.ShouldBindJSON(&questInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	// Create quest
	quest := models.Quest{
		Title: questInput.Title,
		Description: questInput.Description,
		Reward: questInput.Reward,
	}
	savedQuest, err := quest.Save()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{ "data": savedQuest })
}

func FindQuest(context *gin.Context) {
	var quest models.Quest

	if err := models.DB.Where("id = ?", context.Param("albumID")).First(&quest).Error; err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{ "data": quest })
}

func UpdateQuest(context *gin.Context) {
	var quest models.Quest
	if err := models.DB.Where("id = ?", context.Param("albumID")).First(&quest).Error; err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
		return
	}

	// Validate input
	var updateQuestInput models.UpdateQuestInput
	if err := context.ShouldBindJSON(&updateQuestInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	models.DB.Model(&quest).Updates(updateQuestInput)
	context.IndentedJSON(http.StatusOK, gin.H{ "data": quest })
}

func DeleteQuest(context *gin.Context) {
  // Get model if exist
  var quest models.Quest
  if err := models.DB.Where("id = ?", context.Param("albumID")).First(&quest).Error; err != nil {
    context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
    return
  }

  models.DB.Delete(&quest)
  context.IndentedJSON(http.StatusOK, gin.H{ "data": "Deleted successfully" })
}
