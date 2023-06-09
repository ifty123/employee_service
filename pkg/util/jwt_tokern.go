package util

import (
	"encoding/json"
	"fmt"
	"latihan_service/internal/dto"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	JWT_SECRET         = []byte(GetEnv("JWT_SECRET", "testsecret"))
	JWT_EXP            = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

func GetTokenString(authHeader string) (string, error) {
	var token string
	if strings.Contains(authHeader, "Bearer") {
		token = strings.Replace(authHeader, "Bearer", "", -1)
		return token, nil
	}
	return "", fmt.Errorf("authorization not found")
}

func CreateJWTClaims(email string, userId, divisionId uint) dto.JWTClaims {
	return dto.JWTClaims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXP)),
		},
	}
}

func CreateJWTToken(claims dto.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	return token.SignedString([]byte(JWT_SECRET))
}

func ParseJWTToken(authHeader string) (*dto.JWTClaims, error) {
	tokenString, err := GetTokenString(authHeader)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("invalid signing method")
		}
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimStr, err := json.Marshal(claims)
		if err != nil {
			return nil, fmt.Errorf("error when marshalling token")
		}

		var customClaims dto.JWTClaims
		if err := json.Unmarshal(claimStr, &customClaims); err != nil {
			return nil, fmt.Errorf("error when unmarshalling token")
		}

		return &customClaims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
