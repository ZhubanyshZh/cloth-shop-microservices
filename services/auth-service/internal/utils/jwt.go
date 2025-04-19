package utils

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte("ACCESS_SECRET")
var refreshSecret = []byte("REFRESH_SECRET")

func GenerateAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

func GenerateRefreshToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
