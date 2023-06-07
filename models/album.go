package models

import (
	"time"
	// "gorm.io/gorm"
)

// album represents data about a record album.
type Album struct {
	ID uint `gorm:"primary_key" json:"id"` 
	Title string `gorm:"size:255;not null;unique" json:"title"`
	Artist string  `gorm:"size:255;not null" json:"artist"`
	Price  float64 `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type CreateAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type UpdateAlbumInput struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (album *Album) Save() (*Album, error) {
	if err := DB.Create(&album).Error; err != nil {
		return &Album{}, err
	}
	return album, nil
}
