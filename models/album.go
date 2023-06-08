package models

import (
	"time"
	"errors"
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

func (album *Album) Update(updateAlbum UpdateAlbumInput, albumID string) (*Album, error) {
	if err := DB.First(&album, albumID).Error; err != nil {
		return &Album{}, err
	}
	if err := DB.Model(&album).Updates(updateAlbum).Error; err != nil {
		return &Album{}, err
	}
	return album, nil
}

func (album *Album) FindAll() []Album {
	var albums []Album
	DB.Find(&albums)
	return albums
}

func (album *Album) Delete(albumID string) (bool, error) {
	if err := DB.First(&album, albumID).Error; err != nil {
		return false, errors.New("record not found")
	}
	DB.Delete(&album)
	return true, nil
}

func (album *Album) FindAlbumByID(albumID string) (*Album, error) {
	if err := DB.First(&album, albumID).Error; err != nil {
		return album, err
	}
	return album, nil
}
