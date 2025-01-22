package models

import "time"

type URL struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Original    string    `json:"original"`
	Shortened   string    `json:"shortened"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessCount int       `json:"access_count"`
}
