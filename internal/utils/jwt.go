package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId uint, secret string) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(secret))

}
