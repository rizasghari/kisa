package models

import "time"

type URL struct {
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      string    `json:"user_id" gorm:"type:uuid;index"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	OriginalURL string    `json:"original_url" gorm:"not null"`
	ShortURL    string    `json:"short_url" gorm:"unique;not null"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	AccessCount int       `json:"-" gorm:"not null;default:0"`
}
