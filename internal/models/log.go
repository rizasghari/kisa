package models

import "time"

type Log struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UrlID      string    `gorm:"type:uuid;index"`
	URL        URL       `gorm:"foreignKey:UrlID"`
	AccessedAt time.Time `gorm:"autoCreateTime"`
	Referrer   string
	UserAgent  string
}
