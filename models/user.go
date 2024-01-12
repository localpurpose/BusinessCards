package models

import "time"

type User struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	TelegramName string    `json:"telegram_name"`
	CreatedAt    time.Time `json:"created_at"`
	LastUpdate   time.Time `json:"last_update"`
}
