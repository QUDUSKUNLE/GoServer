package models

import "time"

type Quest struct {
  ID uint `json:"id" gorm:"primary_key"`  
  Title string `json:"title"`
  Description string `json:"description"`
  Reward int `json:"reward"`
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
