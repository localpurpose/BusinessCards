package models

type UserSettings struct {
	UserID int64  `json:"user_id"`
	Theme  string `json:"theme"`
	// other Premium modes
}
