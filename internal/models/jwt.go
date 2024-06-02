package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
