package util

import (
	"fiber-golang-json-api/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User) (string, error) {
	secret := []byte("your-secret-key")
	method := jwt.SigningMethodHS256
	claims := jwt.MapClaims{"username": user.Username, "userID": user.ID, "exp": time.Now().Add(24 * time.Hour).Unix()}

	token, err := jwt.NewWithClaims(method, claims).SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
