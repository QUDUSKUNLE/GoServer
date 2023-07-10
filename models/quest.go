package models

import (
	"errors"
	"time"
)

type Quest struct {
	ID          uint      `gorm:"primary_key" json:"questID" `
	Title       string    `gorm:"size:255;not null;unique" json:"title"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Reward      int       `gorm:"not null" json:"reward"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateQuestInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Reward      int    `json:"reward" binding:"required"`
}

type UpdateQuestInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Reward      int    `json:"reward"`
}

func (quest *Quest) Save() (*Quest, error) {
	if err := DB.Create(&quest).Error; err != nil {
		return &Quest{}, err
	}
	return quest, nil
}

func (quest *Quest) GetQuests() []Quest {
	quests := []Quest{}
	DB.Find(&quests)
	return quests
}

func (quest *Quest) GetQuest(id string) (*Quest, error) {
	if err := DB.First(&quest, id).Error; err != nil {
		return quest, err
	}
	return quest, nil
}

func (quest *Quest) UpdateQuest(updateQuest UpdateQuestInput, id string) (*Quest, error) {
	if err := DB.First(&quest, id).Error; err != nil {
		return &Quest{}, err
	}
	if err := DB.Model(&quest).Updates(updateQuest).Error; err != nil {
		return &Quest{}, err
	}
	return quest, nil
}

func (quest *Quest) DeleteQuest(id string) (bool, error) {
	if err := DB.First(&quest, id).Error; err != nil {
		return false, errors.New("record not found")
	}
	DB.Delete(&quest)
	return true, nil
}
