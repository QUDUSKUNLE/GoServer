package models

type Book struct {
	ID      uint `json:"id" gorm:"primary-key"`
	Title   string `json:"title"`
	Author  string `json:"author"`
}
