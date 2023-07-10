package models

import (
	"errors"
	"time"
)

// album represents data about a record album.
type Album struct {
	ID        uint      `gorm:"primary_key" json:"ID"`
	Title     string    `gorm:"size:255;not null;unique" json:"Title"`
	Artist    string    `gorm:"size:255;not null" json:"Artist"`
	Price     float64   `gorm:"not null" json:"Price"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type CreateAlbumInput struct {
	Title  string  `json:"Title" binding:"required"`
	Artist string  `json:"Artist" binding:"required"`
	Price  float64 `json:"Price" binding:"required"`
}

type UpdateAlbumInput struct {
	Title  string  `json:"Title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"Price"`
}

func (album *Album) Save() error {
	if err := DB.Create(&album).Error; err != nil {
		return err
	}
	return nil
}

func (album *Album) Update(updateAlbum UpdateAlbumInput, id string) (*Album, error) {
	if err := DB.First(&album, id).Error; err != nil {
		return &Album{}, err
	}
	if err := DB.Model(&album).Updates(updateAlbum).Error; err != nil {
		return &Album{}, err
	}
	return album, nil
}

func (album *Album) FindAll() []Album {
	albums := []Album{}
	DB.Find(&albums)
	return albums
}

func (album *Album) Delete(id string) (bool, error) {
	if err := DB.First(&album, id).Error; err != nil {
		return false, errors.New("record not found")
	}
	DB.Delete(&album)
	return true, nil
}

func (album *Album) FindAlbumByID(id string) (*Album, error) {
	if err := DB.First(&album, id).Error; err != nil {
		return album, err
	}
	return album, nil
}
