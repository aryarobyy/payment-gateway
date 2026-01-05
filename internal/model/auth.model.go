package model

import "github.com/golang-jwt/jwt/v5"

type ClaimsModel struct {
	UserID string `json:"id"`
	Role   Role   `json:"role"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
