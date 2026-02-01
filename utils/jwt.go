package utils

import (
	"blog-platform/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JWTSecret)
}

func ParseToken(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid claim")
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("User id not found")
	}

	return int(userId), nil

}
