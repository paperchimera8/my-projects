package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("my_secret_key")

func CreateToken(userID string) (string, error) {
	// Создаем Claims (данные внутри токена)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
