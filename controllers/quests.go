package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
)

// GET /quests
// GET all quests
func FindQuests(con *gin.Context) {
	var quests []models.Quest
	models.DB.Find(&quests)
	con.IndentedJSON(http.StatusOK, gin.H{ "data": quests })
}

func CreateQuest(con *gin.Context) {
	// Validate quest
	var questInput models.CreateQuestInput
	if err := con.ShouldBindJSON(&questInput); err != nil {
		con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	// Create quest
	quest := models.Quest{ Title: questInput.Title, Description: questInput.Description, Reward: questInput.Reward}
	models.DB.Create(&quest)

	con.IndentedJSON(http.StatusCreated, gin.H{ "data": quest })
}

func FindQuest(con *gin.Context) {
	var quest models.Quest

	if err := models.DB.Where("id = ?", con.Param("id")).First(&quest).Error; err != nil {
		con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
		return
	}

	con.IndentedJSON(http.StatusOK, gin.H{ "data": quest })
}

func UpdateQuest(con *gin.Context) {
	var quest models.Quest
	if err := models.DB.Where("id = ?", con.Param("id")).First(&quest).Error; err != nil {
		con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
		return
	}

	// Validate input
	var updateQuestInput models.UpdateQuestInput
	if err := con.ShouldBindJSON(&updateQuestInput); err != nil {
		con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	models.DB.Model(&quest).Updates(updateQuestInput)
	con.IndentedJSON(http.StatusOK, gin.H{ "data": quest })
}

func DeleteQuest(con *gin.Context) {
  // Get model if exist
  var quest models.Quest
  if err := models.DB.Where("id = ?", con.Param("id")).First(&quest).Error; err != nil {
    con.IndentedJSON(http.StatusBadRequest, gin.H{ "error": "Record not found!" })
    return
  }

  models.DB.Delete(&quest)
  con.IndentedJSON(http.StatusOK, gin.H{ "data": "Deleted successfully" })
}
