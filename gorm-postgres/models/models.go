package models

import "gorm.io/gorm"

// Book model
type Book struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
}
