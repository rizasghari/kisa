package models

import "time"

type URL struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      string    `gorm:"type:uuid;index"`
	User        User      `gorm:"foreignKey:UserID"`
	OriginalURL string    `gorm:"not null"`
	ShortURL    string    `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	AccessCount int       `gorm:"not null;default:0"`
}
