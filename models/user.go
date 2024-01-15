package models

import "time"

type User struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Organization string    `json:"org"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	WebSite      string    `json:"webSite"`
	TelegramName string    `json:"telegram_name"`
	WhatsApp     string    `json:"whatsApp"`
	QrPath       string    `json:"qrPath"`
	CreatedAt    time.Time `json:"created_at"`
	LastUpdate   time.Time `json:"last_update"`
}
