package models

import (
	"errors"
	"time"
)

type Quest struct {
	ID          uint      `gorm:"primary_key" json:"ID" `
	Title       string    `gorm:"size:255;not null;unique" json:"Title"`
	Description string    `gorm:"size:255;not null" json:"Description"`
	Reward      int       `gorm:"not null" json:"Reward"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

type CreateQuestInput struct {
	Title       string `json:"Title" binding:"required"`
	Description string `json:"Description" binding:"required"`
	Reward      int    `json:"Reward" binding:"required"`
}

type UpdateQuestInput struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Reward      int    `json:"Reward"`
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
