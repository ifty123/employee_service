package dto

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
