package models

import "time"

type User struct {
	ID           string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Password     string    `json:"password" gorm:"-"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
}
