package model

import "time"

// type Blog struct {
// 	ID uint `json:"id" gorm:"primaryKey"`

// 	Title string `json:"title" gorm:"not null;column:title;size:255"`

// 	Post string `json:"post" gorm:"not null;column:post;size:255"`
// }

// type BlogRequest struct {
// 	Title string `json:"title" validate:"required,min=3"`
// 	Post  string `json:"post" validate:"required,min=10"`
// }

type Blog struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255;not null"`
	Post      string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
