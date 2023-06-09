package models

import (
	"time"
  "errors"
)

type Quest struct {
  ID uint `gorm:"primary_key" json:"id" `  
  Title string `gorm:"size:255;not null;unique" json:"title"`
  Description string `gorm:"size:255;not null" json:"description"`
  Reward int `gorm:"not null" json:"reward"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type CreateQuestInput struct {
  Title string `json:"title" binding:"required"`
  Description string `json:"description" binding:"required"`
  Reward int `json:"reward" binding:"required"`
}

type UpdateQuestInput struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Reward int `json:"reward"`
}

func (quest *Quest) Save() (*Quest, error) {
  if err := DB.Create(&quest).Error; err != nil {
    return &Quest{}, err
  }
  return quest, nil
}

func (quest *Quest) GetQuests() []Quest {
  var quests []Quest
  DB.Find(&quests)
  return quests
}

func (quest *Quest) GetQuest(questID string) (*Quest, error) {
  if err := DB.First(&quest, questID).Error; err != nil {
    return quest, err
  }
  return quest, nil
}

func (quest *Quest) UpdateQuest(updateQuest UpdateQuestInput, questID string) (*Quest, error) {
	if err := DB.First(&quest, questID).Error; err != nil {
		return &Quest{}, err
	}
	if err := DB.Model(&quest).Updates(updateQuest).Error; err != nil {
		return &Quest{}, err
	}
	return quest, nil
}

func (quest *Quest) DeleteQuest(questID string) (bool, error) {
	if err := DB.First(&quest, questID).Error; err != nil {
		return false, errors.New("record not found")
	}
	DB.Delete(&quest)
	return true, nil
}
