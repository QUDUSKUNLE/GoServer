package models

import (
	"time"
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

func (quest *Quest) AllQuests() []Quest {
  var quests []Quest
  DB.Find(&quests)
  return quests
}
