package model

import "time"

type Blog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at"`
}
