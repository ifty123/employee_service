package dto

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type RegisterEmployeeReq struct {
	Fullname   string `json:"fullname" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Role       string `json:"role" validate:"required"`
	Password   string `json:"password" validate:"required"`
	DivisionID *uint  `json:"division_id" validate:"required"`
}

type EmailAndPasswordReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
